package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingCheckedOutWithPayment ...
type BookingCheckedOutWithPayment struct {
	message.BaseEvent
	Data BookingCheckedOutWithPaymentData
}

// BookingCheckedOutWithPaymentData ...
type BookingCheckedOutWithPaymentData struct {
	Domain  string                                    `json:"domain"`
	Client  BookingCheckedOutWithPaymentClientData    `json:"client"`
	Tickets []*BookingCheckedOutWithPaymentTicketData `json:"tickets"`
	Buyer   BookingCheckedOutWithPaymentBuyerData     `json:"buyer"`
	Payment BookingCheckedOutWithPaymentPaymentData   `json:"payment"`
	Coupon  *BookingCheckedOutWithPaymentCouponData   `json:"coupon,omitempty"`
}

// BookingCheckedOutWithPaymentCouponData ...
type BookingCheckedOutWithPaymentCouponData struct {
	ID           uuid.UUID   `json:"id"`
	Value        int         `json:"value"`
	DiscountType string      `json:"discountType"`
	TicketIDs    []uuid.UUID `json:"ticketIDs"`
}

// BookingCheckedOutWithPaymentClientData ...
type BookingCheckedOutWithPaymentClientData struct {
	IP                      string  `json:"ip"`
	DeviceFingerprint       *string `json:"deviceFingerprint"`
	UserAgent               *string `json:"userAgent"`
	OS                      *string `json:"os"`
	Source                  string  `json:"source"`
	GoogleAnalyticsClientID *string `json:"googleAnalyticsClientID"`
}

// BookingCheckedOutWithPaymentTicketData ...
type BookingCheckedOutWithPaymentTicketData struct {
	ID        uuid.UUID                                 `json:"id"`
	Trip      BookingCheckedOutWithPaymentTripData      `json:"trip"`
	Subtotal  int                                       `json:"subtotal"`
	VAT       int                                       `json:"vat"`
	Passenger BookingCheckedOutWithPaymentPassengerData `json:"passenger"`
	Seat      BookingCheckedOutWithPaymentSeatData      `json:"seat"`
}

// BookingCheckedOutWithPaymentTripData ...
type BookingCheckedOutWithPaymentTripData struct {
	ID           uuid.UUID                                        `json:"id"`
	Segment      BookingCheckedOutWithPaymentSegmentData          `json:"segment"`
	CompanyLine  BookingCheckedOutWithPaymentTripCompanyLineData  `json:"companyLine"`
	RouteService BookingCheckedOutWithPaymentTripRouteServiceData `json:"routeService"`
}

// BookingCheckedOutWithPaymentTripRouteServiceData ...
type BookingCheckedOutWithPaymentTripRouteServiceData struct {
	ID        uuid.UUID `json:"id"`
	VehicleID uuid.UUID `json:"vehicleID"`
	Code      string    `json:"code"`
}

// BookingCheckedOutWithPaymentSegmentData ...
type BookingCheckedOutWithPaymentSegmentData struct {
	ID          uuid.UUID                                    `json:"id"`
	Origin      BookingCheckedOutWithPaymentSegmentPlaceData `json:"origin"`
	Destination BookingCheckedOutWithPaymentSegmentPlaceData `json:"destination"`
	DepartureAt time.Time                                    `json:"departureAt"`
	ArrivalAt   time.Time                                    `json:"arrivalAt"`
}

// BookingCheckedOutWithPaymentSegmentPlaceData ...
type BookingCheckedOutWithPaymentSegmentPlaceData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingCheckedOutWithPaymentTripCompanyLineData ...
type BookingCheckedOutWithPaymentTripCompanyLineData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingCheckedOutWithPaymentPassengerData ...
type BookingCheckedOutWithPaymentPassengerData struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Type      string    `json:"type"`
}

// BookingCheckedOutWithPaymentSeatData ...
type BookingCheckedOutWithPaymentSeatData struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label"`
}

// BookingCheckedOutWithPaymentBuyerData ...
type BookingCheckedOutWithPaymentBuyerData struct {
	FirstName   string                                     `json:"firstName"`
	LastName    string                                     `json:"lastName"`
	Email       string                                     `json:"email"`
	Phone       string                                     `json:"phone"`
	IsPassenger bool                                       `json:"isPassenger"`
	Card        *BookingCheckedOutWithPaymentBuyerCardData `json:"card,omitempty"`
}

// BookingCheckedOutWithPaymentPaymentData ...
type BookingCheckedOutWithPaymentPaymentData struct {
	SourceToken *string `json:"sourceToken,omitempty"`
	Method      string  `json:"method"`
	Processor   *string `json:"processor,omitempty"`
	Amount      int     `json:"amount"`
	Subtotal    int     `json:"subtotal"`
	VAT         int     `json:"vat"`
	Discount    int     `json:"discount"`
}

// BookingCheckedOutWithPaymentBuyerCardData ...
type BookingCheckedOutWithPaymentBuyerCardData struct {
	ExpirationMonth string                                            `json:"expirationMonth"`
	ExpirationYear  string                                            `json:"expirationYear"`
	BinNumber       string                                            `json:"binNumber"`
	Last4           string                                            `json:"last4"`
	Address         *BookingCheckedOutWithPaymentBuyerCardAddressData `json:"address"`
}

// BookingCheckedOutWithPaymentBuyerCardAddressData ...
type BookingCheckedOutWithPaymentBuyerCardAddressData struct {
	Line1       string  `json:"line1"`
	Line2       *string `json:"line2"`
	City        string  `json:"city"`
	StateCode   string  `json:"stateCode"`
	CountryCode string  `json:"countryCode"`
	ZipCode     string  `json:"zipCode"`
}

// BookingCheckedOutWithPaymentFactory ...
func BookingCheckedOutWithPaymentFactory(e dto.Event) (message.Event, error) {
	var bco BookingCheckedOutWithPayment
	bco.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bco.Data); err != nil {
		return nil, err
	}

	return bco, nil
}

// ConflictWith ...
func (BookingCheckedOutWithPayment) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingCheckedOutWithPayment); ok {
		return true
	}

	return false
}
