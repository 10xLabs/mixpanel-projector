package event

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingReturnTripAddedTripData ...
type BookingReturnTripAddedTripData struct {
	ID             uuid.UUID                             `json:"id"`
	Segment        BookingReturnTripAddedSegmentData     `json:"segment"`
	CompanyLine    BookingReturnTripAddedCompanyLineData `json:"companyLine"`
	RouteServiceID uuid.UUID                             `json:"routeServiceID"`
}

// BookingReturnTripAddedSegmentData ...
type BookingReturnTripAddedSegmentData struct {
	ID          uuid.UUID                            `json:"id"`
	Origin      BookingReturnTripAddedPlaceData      `json:"origin"`
	Destination BookingReturnTripAddedPlaceData      `json:"destination"`
	DepartureAt time.Time                            `json:"departureAt"`
	ArrivalAt   time.Time                            `json:"arrivalAt"`
	FareFamily  BookingReturnTripAddedFareFamilyData `json:"fareFamily"`
}

// BookingReturnTripAddedPlaceData ...
type BookingReturnTripAddedPlaceData struct {
	ID   uuid.UUID                      `json:"id"`
	Name string                         `json:"name"`
	City BookingReturnTripAddedCityData `json:"city"`
}

// BookingReturnTripAddedCityData ...
type BookingReturnTripAddedCityData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingReturnTripAddedCompanyLineData ...
type BookingReturnTripAddedCompanyLineData struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// BookingReturnTripAddedFareFamilyData ...
type BookingReturnTripAddedFareFamilyData struct {
	ID   uuid.UUID                      `json:"id"`
	Fare BookingReturnTripAddedFareData `json:"fare"`
}

// BookingReturnTripAddedFareData ...
type BookingReturnTripAddedFareData struct {
	ID           uuid.UUID                         `json:"id"`
	SalesChannel int                               `json:"salesChannel"`
	Prices       []BookingReturnTripAddedPriceData `json:"prices"`
}

// BookingReturnTripAddedPriceData ...
type BookingReturnTripAddedPriceData struct {
	Subtotal      int    `json:"subtotal"`
	VAT           int    `json:"vat"`
	PassengerType string `json:"passengerType"`
}

// BookingReturnTripAddedData ...
type BookingReturnTripAddedData struct {
	Trip BookingReturnTripAddedTripData `json:"trip"`
}

// BookingReturnTripAdded ...
type BookingReturnTripAdded struct {
	message.BaseEvent
	Data BookingReturnTripAddedData
}

// BookingReturnTripAddedFactory ...
func BookingReturnTripAddedFactory(e dto.Event) (message.Event, error) {
	var ta BookingReturnTripAdded
	ta.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &ta.Data); err != nil {
		return nil, err
	}

	return ta, nil
}

// ConflictWith ...
func (BookingReturnTripAdded) ConflictWith(e message.Event) bool {
	if _, ok := e.(BookingReturnTripAdded); ok {
		return true
	}

	return false
}
