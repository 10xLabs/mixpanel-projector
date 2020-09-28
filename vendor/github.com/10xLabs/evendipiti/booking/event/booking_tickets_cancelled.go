package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingTicketsCancelled ...
type BookingTicketsCancelled struct {
	message.BaseEvent
	Data BookingTicketsCancelledData
}

// BookingTicketsCancelledData ...
type BookingTicketsCancelledData struct {
	Tickets []BookingTicketsCancelledTicketData `json:"tickets"`
}

// BookingTicketsCancelledTicketData ...
type BookingTicketsCancelledTicketData struct {
	ID        uuid.UUID                            `json:"id"`
	Trip      BookingTicketsCancelledTripData      `json:"trip"`
	Subtotal  int                                  `json:"subtotal"`
	VAT       int                                  `json:"vat"`
	Passenger BookingTicketsCancelledPassengerData `json:"passenger"`
	Seat      BookingTicketsCancelledSeatData      `json:"seat"`
}

// BookingTicketsCancelledTripData ...
type BookingTicketsCancelledTripData struct {
	ID          uuid.UUID                                  `json:"id"`
	Segment     BookingTicketsCancelledSegmentData         `json:"segment"`
	CompanyLine BookingTicketsCancelledTripCompanyLineData `json:"companyLine"`
}

// BookingTicketsCancelledSegmentData ...
type BookingTicketsCancelledSegmentData struct {
	ID          uuid.UUID                               `json:"id"`
	Origin      BookingTicketsCancelledSegmentPlaceData `json:"origin"`
	Destination BookingTicketsCancelledSegmentPlaceData `json:"destination"`
	DepartureAt time.Time                               `json:"departureAt"`
	ArrivalAt   time.Time                               `json:"arrivalAt"`
}

// BookingTicketsCancelledSegmentPlaceData ...
type BookingTicketsCancelledSegmentPlaceData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingTicketsCancelledTripCompanyLineData ...
type BookingTicketsCancelledTripCompanyLineData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingTicketsCancelledPassengerData ...
type BookingTicketsCancelledPassengerData struct {
	ID        uuid.UUID `json:"id"`
	FirstName *string   `json:"firstName"`
	LastName  *string   `json:"lastName"`
	Type      string    `json:"type"`
}

// BookingTicketsCancelledSeatData ...
type BookingTicketsCancelledSeatData struct {
	ID uuid.UUID `json:"id"`
}

// BookingTicketsCancelledFactory ...
func BookingTicketsCancelledFactory(e dto.Event) (message.Event, error) {
	var evt BookingTicketsCancelled
	evt.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &evt.Data); err != nil {
		return nil, err
	}

	return evt, nil
}

// ConflictWith ...
func (BookingTicketsCancelled) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingTicketsCancelled); ok {
		return true
	}

	return false
}
