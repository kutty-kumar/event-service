package main

import (
	. "event-service/pkg/pb"
	"event-service/pkg/svc"
	"github.com/golang/protobuf/ptypes"
)

func main() {
	producerConfig := map[string]string{"bootstrap.servers": "localhost:9092"}
	kafkaProducer := svc.NewKafkaProducer([]string{"doctor-created"}, producerConfig)
	event := Event{EntityType: "doctor-created", EntityId: "1", Checksum: "doctor_created_1_checksum",
		CreatedAt: ptypes.TimestampNow()}
	kafkaProducer.Publish(event)
	kafkaProducer.Close()
}
