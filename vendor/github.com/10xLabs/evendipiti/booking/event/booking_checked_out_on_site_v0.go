package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingCheckedOutOnSiteV0 ...
type BookingCheckedOutOnSiteV0 struct {
	message.BaseEvent
	Data BookingCheckedOutOnSiteV0Data
}

// BookingCheckedOutOnSiteV0Data ...
type BookingCheckedOutOnSiteV0Data struct {
	Domain  string                                 `json:"domain"`
	SiteID  uuid.UUID                              `json:"siteID"`
	Tickets []*BookingCheckedOutOnSiteV0TicketData `json:"tickets"`
	Buyer   BookingCheckedOutOnSiteV0BuyerData     `json:"buyer"`
	Payment BookingCheckedOutOnSiteV0PaymentData   `json:"payment"`
}

// BookingCheckedOutOnSiteV0TicketData ...
type BookingCheckedOutOnSiteV0TicketData struct {
	ID        uuid.UUID                              `json:"id"`
	Trip      BookingCheckedOutOnSiteV0TripData      `json:"trip"`
	Subtotal  int                                    `json:"subtotal"`
	VAT       int                                    `json:"vat"`
	Passenger BookingCheckedOutOnSiteV0PassengerData `json:"passenger"`
	Seat      BookingCheckedOutOnSiteV0SeatData      `json:"seat"`
}

// BookingCheckedOutOnSiteV0TripData ...
type BookingCheckedOutOnSiteV0TripData struct {
	ID          uuid.UUID                                    `json:"id"`
	Origin      BookingCheckedOutOnSiteV0TripPlaceData       `json:"origin"`
	Destination BookingCheckedOutOnSiteV0TripPlaceData       `json:"destination"`
	CompanyLine BookingCheckedOutOnSiteV0TripCompanyLineData `json:"companyLine"`
	DepartureAt time.Time                                    `json:"departureAt"`
	ArrivalAt   time.Time                                    `json:"arrivalAt"`
}

// BookingCheckedOutOnSiteV0TripPlaceData ...
type BookingCheckedOutOnSiteV0TripPlaceData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingCheckedOutOnSiteV0TripCompanyLineData ...
type BookingCheckedOutOnSiteV0TripCompanyLineData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingCheckedOutOnSiteV0PassengerData ...
type BookingCheckedOutOnSiteV0PassengerData struct {
	ID        uuid.UUID `json:"id"`
	FirstName *string   `json:"firstName,omitempty"`
	LastName  *string   `json:"lastName,omitempty"`
	Type      string    `json:"type"`
}

// BookingCheckedOutOnSiteV0SeatData ...
type BookingCheckedOutOnSiteV0SeatData struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label"`
}

// BookingCheckedOutOnSiteV0BuyerData ...
type BookingCheckedOutOnSiteV0BuyerData struct {
	Email string `json:"email"`
}

// BookingCheckedOutOnSiteV0PaymentData ...
type BookingCheckedOutOnSiteV0PaymentData struct {
	ID       string  `json:"id"`
	Amount   int     `json:"amount"`
	Method   string  `json:"method"`
	AuthCode *string `json:"authCode,omitempty"`
}

// BookingCheckedOutOnSiteV0Factory ...
func BookingCheckedOutOnSiteV0Factory(e dto.Event) (message.Event, error) {
	var bco BookingCheckedOutOnSiteV0
	bco.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bco.Data); err != nil {
		return nil, err
	}

	return bco, nil
}

// ConflictWith ...
func (BookingCheckedOutOnSiteV0) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingCheckedOutOnSiteV0); ok {
		return true
	}

	return false
}
