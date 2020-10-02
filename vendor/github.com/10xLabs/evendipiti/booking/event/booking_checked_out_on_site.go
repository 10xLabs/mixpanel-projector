package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingCheckedOutOnSite ...
type BookingCheckedOutOnSite struct {
	message.BaseEvent
	Data BookingCheckedOutOnSiteData
}

// BookingCheckedOutOnSiteData ...
type BookingCheckedOutOnSiteData struct {
	Domain  string                               `json:"domain"`
	SiteID  uuid.UUID                            `json:"siteID"`
	Tickets []*BookingCheckedOutOnSiteTicketData `json:"tickets"`
	Buyer   BookingCheckedOutOnSiteBuyerData     `json:"buyer"`
	Payment BookingCheckedOutOnSitePaymentData   `json:"payment"`
	Coupon  *BookingCheckedOutOnSiteCouponData   `json:"coupon,omitempty"`
}

// BookingCheckedOutOnSiteCouponData ...
type BookingCheckedOutOnSiteCouponData struct {
	ID           uuid.UUID   `json:"id"`
	Value        int         `json:"value"`
	DiscountType string      `json:"discountType"`
	TicketIDs    []uuid.UUID `json:"ticketIDs"`
}

// BookingCheckedOutOnSiteTicketData ...
type BookingCheckedOutOnSiteTicketData struct {
	ID        uuid.UUID                            `json:"id"`
	Trip      BookingCheckedOutOnSiteTripData      `json:"trip"`
	Subtotal  int                                  `json:"subtotal"`
	VAT       int                                  `json:"vat"`
	Passenger BookingCheckedOutOnSitePassengerData `json:"passenger"`
	Seat      BookingCheckedOutOnSiteSeatData      `json:"seat"`
}

// BookingCheckedOutOnSiteTripData ...
type BookingCheckedOutOnSiteTripData struct {
	ID           uuid.UUID                                   `json:"id"`
	Segment      BookingCheckedOutOnSiteSegmentData          `json:"segment"`
	CompanyLine  BookingCheckedOutOnSiteTripCompanyLineData  `json:"companyLine"`
	RouteService BookingCheckedOutOnSiteTripRouteServiceData `json:"routeService"`
}

// BookingCheckedOutOnSiteTripRouteServiceData ...
type BookingCheckedOutOnSiteTripRouteServiceData struct {
	ID        uuid.UUID `json:"id"`
	VehicleID uuid.UUID `json:"vehicleID"`
	Code      string    `json:"code"`
}

// BookingCheckedOutOnSiteSegmentData ...
type BookingCheckedOutOnSiteSegmentData struct {
	ID          uuid.UUID                               `json:"id"`
	Origin      BookingCheckedOutOnSiteSegmentPlaceData `json:"origin"`
	Destination BookingCheckedOutOnSiteSegmentPlaceData `json:"destination"`
	DepartureAt time.Time                               `json:"departureAt"`
	ArrivalAt   time.Time                               `json:"arrivalAt"`
}

// BookingCheckedOutOnSiteSegmentPlaceData ...
type BookingCheckedOutOnSiteSegmentPlaceData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingCheckedOutOnSiteTripCompanyLineData ...
type BookingCheckedOutOnSiteTripCompanyLineData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingCheckedOutOnSitePassengerData ...
type BookingCheckedOutOnSitePassengerData struct {
	ID        uuid.UUID `json:"id"`
	FirstName *string   `json:"firstName,omitempty"`
	LastName  *string   `json:"lastName,omitempty"`
	Type      string    `json:"type"`
}

// BookingCheckedOutOnSiteSeatData ...
type BookingCheckedOutOnSiteSeatData struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label"`
}

// BookingCheckedOutOnSiteBuyerData ...
type BookingCheckedOutOnSiteBuyerData struct {
	Email string `json:"email"`
}

// BookingCheckedOutOnSitePaymentData ...
type BookingCheckedOutOnSitePaymentData struct {
	ID       string  `json:"id"`
	Amount   int     `json:"amount"`
	Method   string  `json:"method"`
	AuthCode *string `json:"authCode,omitempty"`
	Subtotal int     `json:"subtotal"`
	VAT      int     `json:"vat"`
	Discount int     `json:"discount"`
}

// BookingCheckedOutOnSiteFactory ...
func BookingCheckedOutOnSiteFactory(e dto.Event) (message.Event, error) {
	var bco BookingCheckedOutOnSite
	bco.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bco.Data); err != nil {
		return nil, err
	}

	return bco, nil
}

// ConflictWith ...
func (BookingCheckedOutOnSite) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingCheckedOutOnSite); ok {
		return true
	}

	return false
}
