package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingPaymentFailed ...
type BookingPaymentFailed struct {
	message.BaseEvent
	Data BookingPaymentFailedData
}

// BookingPaymentFailedData ...
type BookingPaymentFailedData struct {
	Payment PaymentFailedData `json:"payment"`
}

// PaymentFailedData ...
type PaymentFailedData struct {
	ID         uuid.UUID `json:"id"`
	Reason     string    `json:"reason"`
	ExternalID *string   `json:"externalID"`
}

// BookingPaymentFailedFactory ...
func BookingPaymentFailedFactory(e dto.Event) (message.Event, error) {
	var evt BookingPaymentFailed
	evt.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &evt.Data); err != nil {
		return nil, err
	}

	return evt, nil
}

// ConflictWith ...
func (BookingPaymentFailed) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingPaymentFailed); ok {
		return true
	}

	return false
}
