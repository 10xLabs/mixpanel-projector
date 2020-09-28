package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingDepartureTripAddedV0TripData ...
type BookingDepartureTripAddedV0TripData struct {
	ID             uuid.UUID                                  `json:"id"`
	Origin         BookingDepartureTripAddedV0PlaceData       `json:"origin"`
	Destination    BookingDepartureTripAddedV0PlaceData       `json:"destination"`
	CompanyLine    BookingDepartureTripAddedV0CompanyLineData `json:"companyLine"`
	DepartureAt    time.Time                                  `json:"departureAt"`
	ArrivalAt      time.Time                                  `json:"arrivalAt"`
	Fares          []*BookingDepartureTripAddedV0FareData     `json:"fares"`
	RouteServiceID uuid.UUID                                  `json:"routeServiceID"`
}

// BookingDepartureTripAddedV0PlaceData ...
type BookingDepartureTripAddedV0PlaceData struct {
	ID   uuid.UUID                           `json:"id"`
	Name string                              `json:"name"`
	City BookingDepartureTripAddedV0CityData `json:"city"`
}

// BookingDepartureTripAddedV0CityData ...
type BookingDepartureTripAddedV0CityData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingDepartureTripAddedV0CompanyLineData ...
type BookingDepartureTripAddedV0CompanyLineData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingDepartureTripAddedV0FareData ...
type BookingDepartureTripAddedV0FareData struct {
	ID            uuid.UUID `json:"id"`
	Subtotal      int       `json:"subtotal"`
	VAT           int       `json:"vat"`
	PassengerType string    `json:"passengerType"`
	DiscountRate  float64   `json:"discountRate"`
}

// BookingDepartureTripAddedV0Data ...
type BookingDepartureTripAddedV0Data struct {
	Trip BookingDepartureTripAddedV0TripData `json:"trip"`
}

// BookingDepartureTripAddedV0 ...
type BookingDepartureTripAddedV0 struct {
	message.BaseEvent
	Data BookingDepartureTripAddedV0Data
}

// BookingDepartureTripAddedV0Factory ...
func BookingDepartureTripAddedV0Factory(e dto.Event) (message.Event, error) {
	var ta BookingDepartureTripAddedV0
	ta.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &ta.Data); err != nil {
		return nil, err
	}

	return ta, nil
}

// ConflictWith ...
func (BookingDepartureTripAddedV0) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingDepartureTripAddedV0); ok {
		return true
	}

	return false
}
