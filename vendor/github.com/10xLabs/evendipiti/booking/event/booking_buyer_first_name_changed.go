package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
)

// BuyerFirstNameChanged ...
type BuyerFirstNameChanged struct {
	FirstName string `json:"firstName"`
}

// BookingBuyerFirstNameChangedData ...
type BookingBuyerFirstNameChangedData struct {
	Buyer BuyerFirstNameChanged `json:"buyer"`
}

// BookingBuyerFirstNameChanged ...
type BookingBuyerFirstNameChanged struct {
	message.BaseEvent
	Data BookingBuyerFirstNameChangedData
}

// BookingBuyerFirstNameChangedFactory ...
func BookingBuyerFirstNameChangedFactory(e dto.Event) (message.Event, error) {
	var bnc BookingBuyerFirstNameChanged
	bnc.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bnc.Data); err != nil {
		return nil, err
	}

	return bnc, nil
}

// ConflictWith ...
func (BookingBuyerFirstNameChanged) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingBuyerFirstNameChanged); ok {
		return true
	}

	return false
}
