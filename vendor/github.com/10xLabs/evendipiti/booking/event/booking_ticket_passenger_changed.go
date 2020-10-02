package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingTicketsPassengerChanged ...
type BookingTicketsPassengerChanged struct {
	message.BaseEvent
	Data BookingTicketsPassengerChangedData
}

// BookingTicketsPassengerChangedData ...
type BookingTicketsPassengerChangedData struct {
	Tickets []BookingTicketsPassengerChangedTicketData `json:"tickets"`
}

// BookingTicketsPassengerChangedTicketData ...
type BookingTicketsPassengerChangedTicketData struct {
	ID        uuid.UUID                                   `json:"id"`
	Passenger BookingTicketsPassengerChangedPassengerData `json:"passenger"`
}

// BookingTicketsPassengerChangedPassengerData ...
type BookingTicketsPassengerChangedPassengerData struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Type      string    `json:"type"`
}

// BookingTicketsPassengerChangedFactory ...
func BookingTicketsPassengerChangedFactory(e dto.Event) (message.Event, error) {
	var bpc BookingTicketsPassengerChanged
	bpc.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bpc.Data); err != nil {
		return nil, err
	}

	return bpc, nil
}

// ConflictWith ...
func (BookingTicketsPassengerChanged) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingTicketsPassengerChanged); ok {
		return true
	}

	return false
}
