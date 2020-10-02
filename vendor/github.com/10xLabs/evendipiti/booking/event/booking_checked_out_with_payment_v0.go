package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingCheckedOutWithPaymentV0 ...
type BookingCheckedOutWithPaymentV0 struct {
	message.BaseEvent
	Data CheckedOutWithPaymentV0Data
}

// CheckedOutWithPaymentV0Data ...
type CheckedOutWithPaymentV0Data struct {
	Domain  string                               `json:"domain"`
	Client  CheckedOutWithPaymentV0ClientData    `json:"client"`
	Tickets []*CheckedOutWithPaymentV0TicketData `json:"tickets"`
	Buyer   CheckedOutWithPaymentV0BuyerData     `json:"buyer"`
	Payment CheckedOutWithPaymentV0PaymentData   `json:"payment"`
}

// CheckedOutWithPaymentV0ClientData ...
type CheckedOutWithPaymentV0ClientData struct {
	IP                      string  `json:"ip"`
	DeviceFingerprint       *string `json:"deviceFingerprint"`
	UserAgent               *string `json:"userAgent"`
	OS                      *string `json:"os"`
	Source                  string  `json:"source"`
	GoogleAnalyticsClientID *string `json:"googleAnalyticsClientID"`
}

// CheckedOutWithPaymentV0TicketData ...
type CheckedOutWithPaymentV0TicketData struct {
	ID        uuid.UUID                            `json:"id"`
	Trip      CheckedOutWithPaymentV0TripData      `json:"trip"`
	Subtotal  int                                  `json:"subtotal"`
	VAT       int                                  `json:"vat"`
	Passenger CheckedOutWithPaymentV0PassengerData `json:"passenger"`
	Seat      CheckedOutWithPaymentV0SeatData      `json:"seat"`
}

// CheckedOutWithPaymentV0TripData ...
type CheckedOutWithPaymentV0TripData struct {
	ID          uuid.UUID                                  `json:"id"`
	Origin      CheckedOutWithPaymentV0TripPlaceData       `json:"origin"`
	Destination CheckedOutWithPaymentV0TripPlaceData       `json:"destination"`
	CompanyLine CheckedOutWithPaymentV0TripCompanyLineData `json:"companyLine"`
	DepartureAt time.Time                                  `json:"departureAt"`
	ArrivalAt   time.Time                                  `json:"arrivalAt"`
}

// CheckedOutWithPaymentV0TripPlaceData ...
type CheckedOutWithPaymentV0TripPlaceData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// CheckedOutWithPaymentV0TripCompanyLineData ...
type CheckedOutWithPaymentV0TripCompanyLineData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// CheckedOutWithPaymentV0PassengerData ...
type CheckedOutWithPaymentV0PassengerData struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Type      string    `json:"type"`
}

// CheckedOutWithPaymentV0SeatData ...
type CheckedOutWithPaymentV0SeatData struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label"`
}

// CheckedOutWithPaymentV0BuyerData ...
type CheckedOutWithPaymentV0BuyerData struct {
	FirstName   string                               `json:"firstName"`
	LastName    string                               `json:"lastName"`
	Email       string                               `json:"email"`
	Phone       string                               `json:"phone"`
	IsPassenger bool                                 `json:"isPassenger"`
	Card        CheckedOutWithPaymentV0BuyerCardData `json:"card"`
}

// CheckedOutWithPaymentV0PaymentData ...
type CheckedOutWithPaymentV0PaymentData struct {
	SourceToken string `json:"sourceToken"`
	Method      string `json:"method"`
	Processor   string `json:"processor"`
	Amount      int    `json:"amount"`
}

// CheckedOutWithPaymentV0BuyerCardData ...
type CheckedOutWithPaymentV0BuyerCardData struct {
	ExpirationMonth string                                       `json:"expirationMonth"`
	ExpirationYear  string                                       `json:"expirationYear"`
	BinNumber       string                                       `json:"binNumber"`
	Last4           string                                       `json:"last4"`
	Address         *CheckedOutWithPaymentV0BuyerCardAddressData `json:"address"`
}

// CheckedOutWithPaymentV0BuyerCardAddressData ...
type CheckedOutWithPaymentV0BuyerCardAddressData struct {
	Line1       string  `json:"line1"`
	Line2       *string `json:"line2"`
	City        string  `json:"city"`
	StateCode   string  `json:"stateCode"`
	CountryCode string  `json:"countryCode"`
	ZipCode     string  `json:"zipCode"`
}

// BookingCheckedOutWithPaymentV0Factory ...
func BookingCheckedOutWithPaymentV0Factory(e dto.Event) (message.Event, error) {
	var bco BookingCheckedOutWithPaymentV0
	bco.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bco.Data); err != nil {
		return nil, err
	}

	return bco, nil
}

// ConflictWith ...
func (BookingCheckedOutWithPaymentV0) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingCheckedOutWithPaymentV0); ok {
		return true
	}

	return false
}
