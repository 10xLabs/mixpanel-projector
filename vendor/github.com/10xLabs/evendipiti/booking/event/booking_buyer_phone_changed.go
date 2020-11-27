package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
)

// BuyerPhoneChanged ...
type BuyerPhoneChanged struct {
	Phone string `json:"phone"`
}

// BookingBuyerPhoneChangedData ...
type BookingBuyerPhoneChangedData struct {
	Buyer BuyerPhoneChanged `json:"buyer"`
}

// BookingBuyerPhoneChanged ...
type BookingBuyerPhoneChanged struct {
	message.BaseEvent
	Data BookingBuyerPhoneChangedData
}

// BookingBuyerPhoneChangedFactory ...
func BookingBuyerPhoneChangedFactory(e dto.Event) (message.Event, error) {
	var bpc BookingBuyerPhoneChanged
	bpc.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bpc.Data); err != nil {
		return nil, err
	}

	return bpc, nil
}

// ConflictWith ...
func (BookingBuyerPhoneChanged) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingBuyerPhoneChanged); ok {
		return true
	}

	return false
}
