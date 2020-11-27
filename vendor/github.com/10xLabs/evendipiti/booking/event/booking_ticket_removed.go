package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingTicketRemoved ...
type BookingTicketRemoved struct {
	message.BaseEvent
	Data BookingTicketRemovedData
}

// BookingTicketRemovedData ...
type BookingTicketRemovedData struct {
	TicketID uuid.UUID `json:"ticketID"`
}

// BookingTicketRemovedFactory ...
func BookingTicketRemovedFactory(e dto.Event) (message.Event, error) {
	var bta BookingTicketRemoved
	bta.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bta.Data); err != nil {
		return nil, err
	}

	return bta, nil
}

// ConflictWith ...
func (BookingTicketRemoved) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingTicketRemoved); ok {
		return true
	}

	return false
}
