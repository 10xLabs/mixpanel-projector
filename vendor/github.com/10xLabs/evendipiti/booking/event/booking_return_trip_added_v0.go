package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingReturnTripAddedV0TripData ...
type BookingReturnTripAddedV0TripData struct {
	ID             uuid.UUID                               `json:"id"`
	Origin         BookingReturnTripAddedV0PlaceData       `json:"origin"`
	Destination    BookingReturnTripAddedV0PlaceData       `json:"destination"`
	CompanyLine    BookingReturnTripAddedV0CompanyLineData `json:"companyLine"`
	DepartureAt    time.Time                               `json:"departureAt"`
	ArrivalAt      time.Time                               `json:"arrivalAt"`
	Fares          []*BookingReturnTripAddedV0FareData     `json:"fares"`
	RouteServiceID uuid.UUID                               `json:"routeServiceID"`
}

// BookingReturnTripAddedV0PlaceData ...
type BookingReturnTripAddedV0PlaceData struct {
	ID   uuid.UUID                        `json:"id"`
	Name string                           `json:"name"`
	City BookingReturnTripAddedV0CityData `json:"city"`
}

// BookingReturnTripAddedV0CityData ...
type BookingReturnTripAddedV0CityData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingReturnTripAddedV0CompanyLineData ...
type BookingReturnTripAddedV0CompanyLineData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingReturnTripAddedV0FareData ...
type BookingReturnTripAddedV0FareData struct {
	ID            uuid.UUID `json:"id"`
	Subtotal      int       `json:"subtotal"`
	VAT           int       `json:"vat"`
	PassengerType string    `json:"passengerType"`
	DiscountRate  float64   `json:"discountRate"`
}

// BookingReturnTripAddedV0Data ...
type BookingReturnTripAddedV0Data struct {
	Trip BookingReturnTripAddedV0TripData `json:"trip"`
}

// BookingReturnTripAddedV0 ...
type BookingReturnTripAddedV0 struct {
	message.BaseEvent
	Data BookingReturnTripAddedV0Data
}

// BookingReturnTripAddedV0Factory ...
func BookingReturnTripAddedV0Factory(e dto.Event) (message.Event, error) {
	var ta BookingReturnTripAddedV0
	ta.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &ta.Data); err != nil {
		return nil, err
	}

	return ta, nil
}

// ConflictWith ...
func (BookingReturnTripAddedV0) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingReturnTripAddedV0); ok {
		return true
	}

	return false
}
