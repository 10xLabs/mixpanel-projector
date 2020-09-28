package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingCheckedOutV0 ...
type BookingCheckedOutV0 struct {
	message.BaseEvent
	Data CheckedOutV0Data
}

// CheckedOutV0Data ...
type CheckedOutV0Data struct {
	Domain  string                    `json:"domain"`
	Client  CheckedOutV0ClientData    `json:"client"`
	Tickets []*CheckedOutV0TicketData `json:"tickets"`
	Buyer   CheckedOutV0BuyerData     `json:"buyer"`
	Payment CheckedOutV0PaymentData   `json:"payment"`
}

// CheckedOutV0ClientData ...
type CheckedOutV0ClientData struct {
	IP                      string  `json:"ip"`
	DeviceFingerprint       *string `json:"deviceFingerprint"`
	UserAgent               *string `json:"userAgent"`
	OS                      *string `json:"os"`
	Source                  string  `json:"source"`
	GoogleAnalyticsClientID *string `json:"googleAnalyticsClientID"`
}

// CheckedOutV0TicketData ...
type CheckedOutV0TicketData struct {
	ID        uuid.UUID                 `json:"id"`
	Trip      CheckedOutV0TripData      `json:"trip"`
	Subtotal  int                       `json:"subtotal"`
	VAT       int                       `json:"vat"`
	Passenger CheckedOutV0PassengerData `json:"passenger"`
	Seat      CheckedOutV0SeatData      `json:"seat"`
}

// CheckedOutV0TripData ...
type CheckedOutV0TripData struct {
	ID          uuid.UUID                       `json:"id"`
	Origin      CheckedOutV0TripPlaceData       `json:"origin"`
	Destination CheckedOutV0TripPlaceData       `json:"destination"`
	CompanyLine CheckedOutV0TripCompanyLineData `json:"companyLine"`
	DepartureAt time.Time                       `json:"departureAt"`
	ArrivalAt   time.Time                       `json:"arrivalAt"`
}

// CheckedOutV0TripPlaceData ...
type CheckedOutV0TripPlaceData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// CheckedOutV0TripCompanyLineData ...
type CheckedOutV0TripCompanyLineData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// CheckedOutV0PassengerData ...
type CheckedOutV0PassengerData struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Type      string    `json:"type"`
}

// CheckedOutV0SeatData ...
type CheckedOutV0SeatData struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label"`
}

// CheckedOutV0BuyerData ...
type CheckedOutV0BuyerData struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	IsPassenger bool   `json:"isPassenger"`
}

// CheckedOutV0PaymentData ...
type CheckedOutV0PaymentData struct {
	Amount int `json:"amount"`
}

// BookingCheckedOutV0Factory ...
func BookingCheckedOutV0Factory(e dto.Event) (message.Event, error) {
	var bco BookingCheckedOutV0
	bco.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bco.Data); err != nil {
		return nil, err
	}

	return bco, nil
}

// ConflictWith ...
func (BookingCheckedOutV0) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingCheckedOutV0); ok {
		return true
	}

	return false
}
