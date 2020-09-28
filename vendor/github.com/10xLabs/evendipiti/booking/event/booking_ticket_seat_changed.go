package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingTicketSeatChangedSeatData ...
type BookingTicketSeatChangedSeatData struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label"`
}

// BookingTicketSeatChangedTicketData ...
type BookingTicketSeatChangedTicketData struct {
	ID   uuid.UUID                        `json:"id"`
	Seat BookingTicketSeatChangedSeatData `json:"seat"`
}

// BookingTicketSeatChangedData ...
type BookingTicketSeatChangedData struct {
	Ticket BookingTicketSeatChangedTicketData `json:"ticket"`
}

// BookingTicketSeatChanged ...
type BookingTicketSeatChanged struct {
	message.BaseEvent
	Data BookingTicketSeatChangedData
}

// BookingTicketSeatChangedFactory ...
func BookingTicketSeatChangedFactory(e dto.Event) (message.Event, error) {
	var bta BookingTicketSeatChanged
	bta.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bta.Data); err != nil {
		return nil, err
	}

	return bta, nil
}

// ConflictWith ...
func (BookingTicketSeatChanged) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingTicketSeatChanged); ok {
		return true
	}

	return false
}
