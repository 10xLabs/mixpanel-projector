package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingReset ...
type BookingReset struct {
	message.BaseEvent
	Data BookingResetData
}

// BookingResetData ...
type BookingResetData struct {
	Tickets []*BookingResetTicketData `json:"tickets"`
}

// BookingResetTicketData ...
type BookingResetTicketData struct {
	ID     uuid.UUID            `json:"id"`
	TripID uuid.UUID            `json:"tripID"`
	Seat   BookingResetSeatData `json:"seat"`
}

// BookingResetSeatData ...
type BookingResetSeatData struct {
	ID uuid.UUID `json:"id"`
}

// BookingResetFactory ...
func BookingResetFactory(e dto.Event) (message.Event, error) {
	var evt BookingReset
	evt.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &evt.Data); err != nil {
		return nil, err
	}

	return evt, nil
}

// ConflictWith ...
func (BookingReset) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingReset); ok {
		return true
	}

	return false
}
