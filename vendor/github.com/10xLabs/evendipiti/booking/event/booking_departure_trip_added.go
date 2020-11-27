package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingDepartureTripAddedTripData ...
type BookingDepartureTripAddedTripData struct {
	ID             uuid.UUID                                `json:"id"`
	Segment        BookingDepartureTripAddedSegmentData     `json:"segment"`
	RouteServiceID uuid.UUID                                `json:"routeServiceID"`
	CompanyLine    BookingDepartureTripAddedCompanyLineData `json:"companyLine"`
}

// BookingDepartureTripAddedSegmentData ...
type BookingDepartureTripAddedSegmentData struct {
	ID          uuid.UUID                               `json:"id"`
	Origin      BookingDepartureTripAddedPlaceData      `json:"origin"`
	Destination BookingDepartureTripAddedPlaceData      `json:"destination"`
	DepartureAt time.Time                               `json:"departureAt"`
	ArrivalAt   time.Time                               `json:"arrivalAt"`
	FareFamily  BookingDepartureTripAddedFareFamilyData `json:"fareFamily"`
}

// BookingDepartureTripAddedPlaceData ...
type BookingDepartureTripAddedPlaceData struct {
	ID   uuid.UUID                         `json:"id"`
	Name string                            `json:"name"`
	City BookingDepartureTripAddedCityData `json:"city"`
}

// BookingDepartureTripAddedCityData ...
type BookingDepartureTripAddedCityData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingDepartureTripAddedCompanyLineData ...
type BookingDepartureTripAddedCompanyLineData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingDepartureTripAddedFareFamilyData ...
type BookingDepartureTripAddedFareFamilyData struct {
	ID   uuid.UUID                         `json:"id"`
	Fare BookingDepartureTripAddedFareData `json:"fare"`
}

// BookingDepartureTripAddedFareData ...
type BookingDepartureTripAddedFareData struct {
	ID           uuid.UUID                            `json:"id"`
	SalesChannel int                                  `json:"salesChannel"`
	Prices       []BookingDepartureTripAddedPriceData `json:"prices"`
}

// BookingDepartureTripAddedPriceData ...
type BookingDepartureTripAddedPriceData struct {
	Subtotal      int    `json:"subtotal"`
	VAT           int    `json:"vat"`
	PassengerType string `json:"passengerType"`
}

// BookingDepartureTripAddedData ...
type BookingDepartureTripAddedData struct {
	Trip BookingDepartureTripAddedTripData `json:"trip"`
}

// BookingDepartureTripAdded ...
type BookingDepartureTripAdded struct {
	message.BaseEvent
	Data BookingDepartureTripAddedData
}

// BookingDepartureTripAddedFactory ...
func BookingDepartureTripAddedFactory(e dto.Event) (message.Event, error) {
	var ta BookingDepartureTripAdded
	ta.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &ta.Data); err != nil {
		return nil, err
	}

	return ta, nil
}

// ConflictWith ...
func (BookingDepartureTripAdded) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingDepartureTripAdded); ok {
		return true
	}

	return false
}
