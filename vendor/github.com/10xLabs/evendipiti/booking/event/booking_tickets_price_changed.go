package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingTicketsPriceChanged ...
type BookingTicketsPriceChanged struct {
	message.BaseEvent
	Data BookingTicketsPriceChangedData
}

// BookingTicketsPriceChangedData ...
type BookingTicketsPriceChangedData struct {
	Tickets []BookingTicketsPriceChangedTicketData `json:"tickets"`
}

// BookingTicketsPriceChangedTicketData ...
type BookingTicketsPriceChangedTicketData struct {
	ID       uuid.UUID `json:"id"`
	Subtotal int       `json:"subtotal"`
	VAT      int       `json:"vat"`
}

// BookingTicketsPriceChangedFactory ...
func BookingTicketsPriceChangedFactory(e dto.Event) (message.Event, error) {
	var btpc BookingTicketsPriceChanged
	btpc.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &btpc.Data); err != nil {
		return nil, err
	}

	return btpc, nil
}

// ConflictWith ...
func (BookingTicketsPriceChanged) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingTicketsPriceChanged); ok {
		return true
	}

	return false
}
