package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
)

// BookingTicketsConfirmationFailed ...
type BookingTicketsConfirmationFailed struct {
	message.BaseEvent
	Data BookingTicketsConfirmationFailedData
}

// BookingTicketsConfirmationFailedData ...
type BookingTicketsConfirmationFailedData struct {
	Error BookingTicketsConfirmationFailedDataError `json:"error"`
}

// BookingTicketsConfirmationFailedDataError ...
type BookingTicketsConfirmationFailedDataError struct {
	Code    string  `json:"code"`
	Message *string `json:"message"`
}

// BookingTicketsConfirmationFailedFactory ...
func BookingTicketsConfirmationFailedFactory(e dto.Event) (message.Event, error) {
	var evt BookingTicketsConfirmationFailed
	evt.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &evt.Data); err != nil {
		return nil, err
	}

	return evt, nil
}

// ConflictWith ...
func (BookingTicketsConfirmationFailed) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingTicketsConfirmationFailed); ok {
		return true
	}

	return false
}
