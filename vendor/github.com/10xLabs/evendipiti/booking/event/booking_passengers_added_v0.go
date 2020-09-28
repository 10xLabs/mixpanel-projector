package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingPassengersAddedV0PassengerData ...
type BookingPassengersAddedV0PassengerData struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Type      string    `json:"type"`
}

// BookingPassengersAddedV0Data ...
type BookingPassengersAddedV0Data struct {
	Passengers []*BookingPassengersAddedV0PassengerData `json:"passengers"`
	TripIDs    []uuid.UUID                              `json:"tripIDs"`
}

// BookingPassengersAddedV0 ...
type BookingPassengersAddedV0 struct {
	message.BaseEvent
	Data BookingPassengersAddedV0Data
}

// BookingPassengersAddedV0Factory ...
func BookingPassengersAddedV0Factory(e dto.Event) (message.Event, error) {
	var ta BookingPassengersAddedV0
	ta.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &ta.Data); err != nil {
		return nil, err
	}

	return ta, nil
}

// ConflictWith ...
func (BookingPassengersAddedV0) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingPassengersAddedV0); ok {
		return true
	}

	return false
}
