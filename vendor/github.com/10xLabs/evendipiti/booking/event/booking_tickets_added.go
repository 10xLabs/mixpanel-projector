package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingTicketsAddedPassengerData ...
type BookingTicketsAddedPassengerData struct {
	ID        uuid.UUID `json:"id"`
	FirstName *string   `json:"firstName,omitempty"`
	LastName  *string   `json:"lastName,omitempty"`
	Type      string    `json:"type"`
}

// BookingTicketsAddedSeatData ...
type BookingTicketsAddedSeatData struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label"`
}

// BookingTicketsAddedTripData ...
type BookingTicketsAddedTripData struct {
	ID        uuid.UUID `json:"id"`
	SegmentID uuid.UUID `json:"segmentID"`
}

// BookingTicketsAddedTicketData ...
type BookingTicketsAddedTicketData struct {
	ID        uuid.UUID                        `json:"id"`
	Trip      BookingTicketsAddedTripData      `json:"trip"`
	Subtotal  int                              `json:"subtotal"`
	VAT       int                              `json:"vat"`
	Passenger BookingTicketsAddedPassengerData `json:"passenger"`
	ExpiresAt time.Time                        `json:"expiresAt"`
	Seat      BookingTicketsAddedSeatData      `json:"seat"`
}

// BookingTicketsAddedData ...
type BookingTicketsAddedData struct {
	Tickets []BookingTicketsAddedTicketData `json:"tickets"`
}

// BookingTicketsAdded ...
type BookingTicketsAdded struct {
	message.BaseEvent
	Data BookingTicketsAddedData
}

// BookingTicketsAddedFactory ...
func BookingTicketsAddedFactory(e dto.Event) (message.Event, error) {
	var bta BookingTicketsAdded
	bta.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bta.Data); err != nil {
		return nil, err
	}

	return bta, nil
}

// ConflictWith ...
func (BookingTicketsAdded) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingTicketsAdded); ok {
		return true
	}

	return false
}
