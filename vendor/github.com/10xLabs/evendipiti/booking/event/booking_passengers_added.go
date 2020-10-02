package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingPassengersAddedPassengerData ...
type BookingPassengersAddedPassengerData struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Type      string    `json:"type"`
}

// BookingPassengersAddedData ...
type BookingPassengersAddedData struct {
	Passengers []*BookingPassengersAddedPassengerData `json:"passengers"`
	Trips      []BookingPassengersAddedDataTrip       `json:"trips"`
}

// BookingPassengersAddedDataTrip ...
type BookingPassengersAddedDataTrip struct {
	ID        uuid.UUID `json:"id"`
	SegmentID uuid.UUID `json:"segmentID"`
}

// BookingPassengersAdded ...
type BookingPassengersAdded struct {
	message.BaseEvent
	Data BookingPassengersAddedData
}

// BookingPassengersAddedFactory ...
func BookingPassengersAddedFactory(e dto.Event) (message.Event, error) {
	var ta BookingPassengersAdded
	ta.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &ta.Data); err != nil {
		return nil, err
	}

	return ta, nil
}

// ConflictWith ...
func (BookingPassengersAdded) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingPassengersAdded); ok {
		return true
	}

	return false
}
