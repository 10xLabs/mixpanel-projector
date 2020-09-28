package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingCheckedOut ...
type BookingCheckedOut struct {
	message.BaseEvent
	Data BookingCheckedOutData
}

// BookingCheckedOutData ...
type BookingCheckedOutData struct {
	Domain  string                         `json:"domain"`
	Client  BookingCheckedOutClientData    `json:"client"`
	Tickets []*BookingCheckedOutTicketData `json:"tickets"`
	Buyer   BookingCheckedOutBuyerData     `json:"buyer"`
	Payment BookingCheckedOutPaymentData   `json:"payment"`
	Coupon  *BookingCheckedOutCouponData   `json:"coupon,omitempty"`
}

// BookingCheckedOutClientData ...
type BookingCheckedOutClientData struct {
	IP                      string  `json:"ip"`
	DeviceFingerprint       *string `json:"deviceFingerprint"`
	UserAgent               *string `json:"userAgent"`
	OS                      *string `json:"os"`
	Source                  string  `json:"source"`
	GoogleAnalyticsClientID *string `json:"googleAnalyticsClientID"`
}

// BookingCheckedOutTicketData ...
type BookingCheckedOutTicketData struct {
	ID        uuid.UUID                      `json:"id"`
	Trip      BookingCheckedOutTripData      `json:"trip"`
	Subtotal  int                            `json:"subtotal"`
	VAT       int                            `json:"vat"`
	Passenger BookingCheckedOutPassengerData `json:"passenger"`
	Seat      BookingCheckedOutSeatData      `json:"seat"`
}

// BookingCheckedOutTripData ...
type BookingCheckedOutTripData struct {
	ID           uuid.UUID                             `json:"id"`
	Segment      BookingCheckedOutSegmentData          `json:"segment"`
	CompanyLine  BookingCheckedOutTripCompanyLineData  `json:"companyLine"`
	RouteService BookingCheckedOutTripRouteServiceData `json:"routeService"`
}

// BookingCheckedOutTripRouteServiceData ...
type BookingCheckedOutTripRouteServiceData struct {
	ID        uuid.UUID `json:"id"`
	VehicleID uuid.UUID `json:"vehicleID"`
	Code      string    `json:"code"`
}

// BookingCheckedOutSegmentData ...
type BookingCheckedOutSegmentData struct {
	ID          uuid.UUID                         `json:"id"`
	Origin      BookingCheckedOutSegmentPlaceData `json:"origin"`
	Destination BookingCheckedOutSegmentPlaceData `json:"destination"`
	DepartureAt time.Time                         `json:"departureAt"`
	ArrivalAt   time.Time                         `json:"arrivalAt"`
}

// BookingCheckedOutSegmentPlaceData ...
type BookingCheckedOutSegmentPlaceData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingCheckedOutTripCompanyLineData ...
type BookingCheckedOutTripCompanyLineData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingCheckedOutPassengerData ...
type BookingCheckedOutPassengerData struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Type      string    `json:"type"`
}

// BookingCheckedOutSeatData ...
type BookingCheckedOutSeatData struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label"`
}

// BookingCheckedOutBuyerData ...
type BookingCheckedOutBuyerData struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	IsPassenger bool   `json:"isPassenger"`
}

// BookingCheckedOutPaymentData ...
type BookingCheckedOutPaymentData struct {
	Amount   int `json:"amount"`
	Subtotal int `json:"subtotal"`
	VAT      int `json:"vat"`
	Discount int `json:"discount"`
}

// BookingCheckedOutCouponData ...
type BookingCheckedOutCouponData struct {
	ID           uuid.UUID   `json:"id"`
	Value        int         `json:"value"`
	DiscountType string      `json:"discountType"`
	TicketIDs    []uuid.UUID `json:"ticketIDs"`
}

// BookingCheckedOutFactory ...
func BookingCheckedOutFactory(e dto.Event) (message.Event, error) {
	var bco BookingCheckedOut
	bco.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bco.Data); err != nil {
		return nil, err
	}

	return bco, nil
}

// ConflictWith ...
func (BookingCheckedOut) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingCheckedOut); ok {
		return true
	}

	return false
}
