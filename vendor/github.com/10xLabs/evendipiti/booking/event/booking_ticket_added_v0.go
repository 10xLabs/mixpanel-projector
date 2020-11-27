package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingTicketAddedV0PassengerData ...
type BookingTicketAddedV0PassengerData struct {
	ID        uuid.UUID `json:"id"`
	FirstName *string   `json:"firstName,omitempty"`
	LastName  *string   `json:"lastName,omitempty"`
	Type      string    `json:"type"`
}

// BookingTicketAddedV0SeatData ...
type BookingTicketAddedV0SeatData struct {
	ID        uuid.UUID `json:"id"`
	Label     string    `json:"label"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// BookingTicketAddedV0TicketData ...
type BookingTicketAddedV0TicketData struct {
	ID        uuid.UUID                         `json:"id"`
	TripID    uuid.UUID                         `json:"tripID"`
	Subtotal  int                               `json:"subtotal"`
	VAT       int                               `json:"vat"`
	Passenger BookingTicketAddedV0PassengerData `json:"passenger"`
	Seat      BookingTicketAddedV0SeatData      `json:"seat"`
}

// BookingTicketAddedV0Data ...
type BookingTicketAddedV0Data struct {
	Ticket BookingTicketAddedV0TicketData `json:"ticket"`
}

// BookingTicketAddedV0 ...
type BookingTicketAddedV0 struct {
	message.BaseEvent
	Data BookingTicketAddedV0Data
}

// BookingTicketAddedV0Factory ...
func BookingTicketAddedV0Factory(e dto.Event) (message.Event, error) {
	var bta BookingTicketAddedV0
	bta.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bta.Data); err != nil {
		return nil, err
	}

	return bta, nil
}

// ConflictWith ...
func (BookingTicketAddedV0) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingTicketAddedV0); ok {
		return true
	}

	return false
}
