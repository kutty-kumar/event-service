package svc

import (
	"event-service/pkg/pb"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"sync"
)

type KafkaPublisher struct {
	done          chan bool
	eventChannels []string
	producer      *kafka.Producer
	wg            *sync.WaitGroup
}

func newKafkaProducerConfig(config map[string]string) *kafka.ConfigMap {
	kafkaConfigMap := &kafka.ConfigMap{}
	for key, value := range config {
		err := kafkaConfigMap.SetKey(key, value)
		if err != nil {
			log.Println(fmt.Sprintf("An error %v occurred while setting key %v against %v", err, key, value))
		}
	}
	return kafkaConfigMap
}

func NewKafkaProducer(channels []string, config map[string]string) *KafkaPublisher {
	var wg sync.WaitGroup
	kafkaConfigMap := newKafkaProducerConfig(config)
	producer, err := kafka.NewProducer(kafkaConfigMap)
	if err != nil {
		log.Fatalf(fmt.Sprintf("An error %v occurred while starting kafka producer.", err))
	}
	done := make(chan bool)
	return &KafkaPublisher{done: done, eventChannels: channels, wg: &wg, producer: producer}
}

func (kP *KafkaPublisher) Publish(event eventService.Event) {
	log.Println("Calling publish")
	kP.wg.Add(1)
	go func() {
		defer kP.wg.Done()
		eventBytes, err := proto.Marshal(&event)
		if err != nil {
			log.Println(fmt.Sprintf("An error %v occurred while marshalling %v into bytes.", err, event))
			return
		}
		kP.producer.ProduceChannel() <- &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &event.EntityType, Partition: kafka.PartitionAny},
			Value:          eventBytes,
		}
		log.Println(fmt.Printf("Produced %v to kafka", event))
	}()

}

func (kP *KafkaPublisher) PublishAsync(event eventService.Event) {
	eventBytes, err := proto.Marshal(&event)
	if err != nil {
		log.Println(fmt.Sprintf("An error %v occurred while marshalling %v into bytes.", err, event))
		return
	}
	kP.producer.ProduceChannel() <- &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &event.EntityType, Partition: kafka.PartitionAny},
		Value:          eventBytes,
	}
}

func (kP *KafkaPublisher) Flush() {
	kP.producer.Flush(15 * 1000)
}

func (kP *KafkaPublisher) Close() {
	kP.wg.Wait()
	kP.producer.Close()
}
