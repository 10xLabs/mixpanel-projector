package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingTicketCheckedIn ...
type BookingTicketCheckedIn struct {
	message.BaseEvent
	Data BookingTicketCheckedInData
}

// BookingTicketCheckedInData ...
type BookingTicketCheckedInData struct {
	TicketID uuid.UUID
}

// BookingTicketCheckedInFactory ...
func BookingTicketCheckedInFactory(e dto.Event) (message.Event, error) {
	var evt BookingTicketCheckedIn
	evt.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &evt.Data); err != nil {
		return nil, err
	}

	return evt, nil
}

// ConflictWith ...
func (BookingTicketCheckedIn) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingTicketCheckedIn); ok {
		return true
	}

	return false
}
