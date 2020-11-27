package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
)

// BookingBuyerMarkedAsPassengerData ...
type BookingBuyerMarkedAsPassengerData struct{}

// BookingBuyerMarkedAsPassenger ...
type BookingBuyerMarkedAsPassenger struct {
	message.BaseEvent
	Data BookingBuyerMarkedAsPassengerData
}

// BookingBuyerMarkedAsPassengerFactory ...
func BookingBuyerMarkedAsPassengerFactory(e dto.Event) (message.Event, error) {
	var bmp BookingBuyerMarkedAsPassenger
	bmp.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bmp.Data); err != nil {
		return nil, err
	}

	return bmp, nil
}

// ConflictWith ...
func (BookingBuyerMarkedAsPassenger) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingBuyerMarkedAsPassenger); ok {
		return true
	}

	return false
}
