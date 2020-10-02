package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
)

// BuyerEmailChanged ...
type BuyerEmailChanged struct {
	Email string `json:"email"`
}

// BookingBuyerEmailChangedData ...
type BookingBuyerEmailChangedData struct {
	Buyer BuyerEmailChanged `json:"buyer"`
}

// BookingBuyerEmailChanged ...
type BookingBuyerEmailChanged struct {
	message.BaseEvent
	Data BookingBuyerEmailChangedData
}

// BookingBuyerEmailChangedFactory ...
func BookingBuyerEmailChangedFactory(e dto.Event) (message.Event, error) {
	var bec BookingBuyerEmailChanged
	bec.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bec.Data); err != nil {
		return nil, err
	}

	return bec, nil
}

// ConflictWith ...
func (BookingBuyerEmailChanged) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingBuyerEmailChanged); ok {
		return true
	}

	return false
}
