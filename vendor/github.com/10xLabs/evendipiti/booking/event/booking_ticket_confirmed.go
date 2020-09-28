package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingTicketsConfirmed ...
type BookingTicketsConfirmed struct {
	message.BaseEvent
	Data BookingTicketsConfirmedData
}

// BookingTicketsConfirmedData ...
type BookingTicketsConfirmedData struct {
	Tickets []BookingTicketsConfirmedTicketData
}

// BookingTicketsConfirmedTicketData ...
type BookingTicketsConfirmedTicketData struct {
	SeatID uuid.UUID
}

// BookingTicketsConfirmedFactory ...
func BookingTicketsConfirmedFactory(e dto.Event) (message.Event, error) {
	var btc BookingTicketsConfirmed
	btc.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &btc.Data); err != nil {
		return nil, err
	}

	return btc, nil
}

// ConflictWith ...
func (BookingTicketsConfirmed) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingTicketsConfirmed); ok {
		return true
	}

	return false
}
