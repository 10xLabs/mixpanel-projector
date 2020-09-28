package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingPaid ...
type BookingPaid struct {
	message.BaseEvent
	Data BookingPaidData
}

// BookingPaidData ...
type BookingPaidData struct {
	Tickets []*BookingPaidTicketData `json:"tickets"`
	Payment BookingPaidPaymentData   `json:"payment"`
}

// BookingPaidTicketData ...
type BookingPaidTicketData struct {
	ID        uuid.UUID                `json:"id"`
	Trip      BookingPaidTripData      `json:"trip"`
	Subtotal  int                      `json:"subtotal"`
	VAT       int                      `json:"vat"`
	Passenger BookingPaidPassengerData `json:"passenger"`
	Seat      BookingPaidSeatData      `json:"seat"`
}

// BookingPaidTripData ...
type BookingPaidTripData struct {
	ID          uuid.UUID                      `json:"id"`
	Segment     BookingPaidSegmentData         `json:"segment"`
	CompanyLine BookingPaidTripCompanyLineData `json:"companyLine"`
}

// BookingPaidSegmentData ...
type BookingPaidSegmentData struct {
	ID          uuid.UUID                   `json:"id"`
	Origin      BookingPaidSegmentPlaceData `json:"origin"`
	Destination BookingPaidSegmentPlaceData `json:"destination"`
	DepartureAt time.Time                   `json:"departureAt"`
	ArrivalAt   time.Time                   `json:"arrivalAt"`
}

// BookingPaidSegmentPlaceData ...
type BookingPaidSegmentPlaceData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingPaidTripCompanyLineData ...
type BookingPaidTripCompanyLineData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingPaidPassengerData ...
type BookingPaidPassengerData struct {
	ID        uuid.UUID `json:"id"`
	FirstName *string   `json:"firstName"`
	LastName  *string   `json:"lastName"`
	Type      string    `json:"type"`
}

// BookingPaidSeatData ...
type BookingPaidSeatData struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label"`
}

// BookingPaidPaymentData ...
type BookingPaidPaymentData struct {
	ID            uuid.UUID `json:"id"`
	AuthCode      string    `json:"authCode"`
	ExternalID    string    `json:"externalID"`
	Fee           int32     `json:"fee"`
	TicketsAmount int       `json:"ticketsAmount"`
}

// BookingPaidFactory ...
func BookingPaidFactory(e dto.Event) (message.Event, error) {
	var b BookingPaid
	b.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &b.Data); err != nil {
		return nil, err
	}

	return b, nil
}

// ConflictWith ...
func (BookingPaid) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingPaid); ok {
		return true
	}

	return false
}
