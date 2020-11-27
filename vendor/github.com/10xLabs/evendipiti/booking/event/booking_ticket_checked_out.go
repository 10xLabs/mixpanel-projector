package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingTicketCheckedOut ...
type BookingTicketCheckedOut struct {
	message.BaseEvent
	Data BookingTicketCheckedOutData
}

// BookingTicketCheckedOutData ...
type BookingTicketCheckedOutData struct {
	TicketID uuid.UUID
}

// BookingTicketCheckedOutFactory ...
func BookingTicketCheckedOutFactory(e dto.Event) (message.Event, error) {
	var evt BookingTicketCheckedOut
	evt.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &evt.Data); err != nil {
		return nil, err
	}

	return evt, nil
}

// ConflictWith ...
func (BookingTicketCheckedOut) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingTicketCheckedOut); ok {
		return true
	}

	return false
}
