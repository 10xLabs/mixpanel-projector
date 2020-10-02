package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
)

// BookingBuyerUnmarkedAsPassengerData ...
type BookingBuyerUnmarkedAsPassengerData struct{}

// BookingBuyerUnmarkedAsPassenger ...
type BookingBuyerUnmarkedAsPassenger struct {
	message.BaseEvent
	Data BookingBuyerUnmarkedAsPassengerData
}

// BookingBuyerUnmarkedAsPassengerFactory ...
func BookingBuyerUnmarkedAsPassengerFactory(e dto.Event) (message.Event, error) {
	var bup BookingBuyerUnmarkedAsPassenger
	bup.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bup.Data); err != nil {
		return nil, err
	}

	return bup, nil
}

// ConflictWith ...
func (BookingBuyerUnmarkedAsPassenger) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingBuyerUnmarkedAsPassenger); ok {
		return true
	}

	return false
}
