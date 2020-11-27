package event

import (
	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
)

// BookingConfirmed ...
type BookingConfirmed struct {
	message.BaseEvent
}

// BookingConfirmedFactory ...
func BookingConfirmedFactory(e dto.Event) (message.Event, error) {
	return BookingConfirmed{
		BaseEvent: message.NewBaseEvent(e),
	}, nil
}

// ConflictWith ...
func (BookingConfirmed) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingConfirmed); ok {
		return true
	}

	return false
}
