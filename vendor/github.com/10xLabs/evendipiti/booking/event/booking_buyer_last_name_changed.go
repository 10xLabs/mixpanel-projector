package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
)

// BuyerLastNameChanged ...
type BuyerLastNameChanged struct {
	LastName string `json:"lastName"`
}

// BookingBuyerLastNameChangedData ...
type BookingBuyerLastNameChangedData struct {
	Buyer BuyerLastNameChanged `json:"buyer"`
}

// BookingBuyerLastNameChanged ...
type BookingBuyerLastNameChanged struct {
	message.BaseEvent
	Data BookingBuyerLastNameChangedData
}

// BookingBuyerLastNameChangedFactory ...
func BookingBuyerLastNameChangedFactory(e dto.Event) (message.Event, error) {
	var bnc BookingBuyerLastNameChanged
	bnc.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bnc.Data); err != nil {
		return nil, err
	}

	return bnc, nil
}

// ConflictWith ...
func (BookingBuyerLastNameChanged) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingBuyerLastNameChanged); ok {
		return true
	}

	return false
}
