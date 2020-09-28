package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
)

// BookingTicketsReservationFailed ...
type BookingTicketsReservationFailed struct {
	message.BaseEvent
	Data BookingTicketsReservationFailedData
}

// BookingTicketsReservationFailedData ...
type BookingTicketsReservationFailedData struct {
	Error BookingTicketsReservationFailedDataError `json:"error"`
}

// BookingTicketsReservationFailedDataError ...
type BookingTicketsReservationFailedDataError struct {
	Code    string  `json:"code"`
	Message *string `json:"message"`
}

// BookingTicketsReservationFailedFactory ...
func BookingTicketsReservationFailedFactory(e dto.Event) (message.Event, error) {
	var evt BookingTicketsReservationFailed
	evt.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &evt.Data); err != nil {
		return nil, err
	}

	return evt, nil
}

// ConflictWith ...
func (BookingTicketsReservationFailed) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingTicketsReservationFailed); ok {
		return true
	}

	return false
}
