package svc

import "event-service/pkg/pb"

type Publisher interface {
	Publish(event eventService.Event)
	PublishAsync(event eventService.Event)
	Flush()
	Close()
}
