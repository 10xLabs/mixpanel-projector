package message

import (
	"fmt"
	"log"

	"github.com/10xLabs/comandler/message/dto"
)

type eventFactory func(dto.Event) (Event, error)

var eventFactoryRegister = map[string]eventFactory{}

// RegisterEventFactory ...
func RegisterEventFactory(eventType string, f eventFactory) {
	if f == nil {
		log.Panicf("event factory for %s is nil", eventType)
	}

	if _, ok := eventFactoryRegister[eventType]; ok {
		fmt.Printf("event factory for %s will be overwritten", eventType)
	}

	eventFactoryRegister[eventType] = f
}

// CreateEvent ...
func CreateEvent(e dto.Event) (Event, error) {
	factory := eventFactoryRegister[e.Type]
	if factory == nil {
		return nil, ErrUnregisteredEventFactory
	}

	return factory(e)
}
