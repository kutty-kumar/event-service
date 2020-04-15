package main

import (
	"event-service/pkg/pb"
	"event-service/pkg/svc"
	"fmt"
	"log"
	"sync"
)

func main() {
	eventConsumerMapping := map[string]svc.EventConsumer{"doctor-created": func(event eventService.Event) {
		log.Println(fmt.Sprintf("Got event %v from kafka", event))
	}}
	consumerConfig := map[string]interface{}{
		"bootstrap.servers":               "localhost:9092",
		"group.id":                        "catalog-2",
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"enable.partition.eof":            true,
		"auto.offset.reset":               "earliest",
	}
	var wg sync.WaitGroup
	kafkaConsumer := svc.NewKafkaConsumer([]string{"doctor-created"}, consumerConfig, eventConsumerMapping)
	kafkaConsumer.Consume(&wg)
	wg.Wait()
	kafkaConsumer.Close()
}
