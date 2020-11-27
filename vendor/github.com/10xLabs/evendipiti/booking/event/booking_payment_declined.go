package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingPaymentDeclined ...
type BookingPaymentDeclined struct {
	message.BaseEvent
	Data BookingPaymentDeclinedData
}

// BookingPaymentDeclinedData ...
type BookingPaymentDeclinedData struct {
	Payment PaymentDeclinedData `json:"payment"`
}

// PaymentDeclinedData ...
type PaymentDeclinedData struct {
	ID         uuid.UUID `json:"id"`
	Reason     string    `json:"reason"`
	ExternalID string    `json:"externalID"`
}

// BookingPaymentDeclinedFactory ...
func BookingPaymentDeclinedFactory(e dto.Event) (message.Event, error) {
	var evt BookingPaymentDeclined
	evt.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &evt.Data); err != nil {
		return nil, err
	}

	return evt, nil
}

// ConflictWith ...
func (BookingPaymentDeclined) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingPaymentDeclined); ok {
		return true
	}

	return false
}
