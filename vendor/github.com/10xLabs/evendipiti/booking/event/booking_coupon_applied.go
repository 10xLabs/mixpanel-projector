package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingCouponApplied ...
type BookingCouponApplied struct {
	message.BaseEvent
	Data BookingCouponAppliedData
}

// BookingCouponAppliedData ...
type BookingCouponAppliedData struct {
	Coupon  BookingCouponAppliedCouponData  `json:"coupon"`
	Payment BookingCouponAppliedPaymentData `json:"payment"`
}

// BookingCouponAppliedCouponData ...
type BookingCouponAppliedCouponData struct {
	ID           uuid.UUID   `json:"id"`
	Value        int         `json:"value"`
	DiscountType string      `json:"discountType"`
	TicketIDs    []uuid.UUID `json:"ticketIDs"`
}

// BookingCouponAppliedPaymentData ...
type BookingCouponAppliedPaymentData struct {
	Subtotal int `json:"subtotal"`
	VAT      int `json:"vat"`
	Discount int `json:"discount"`
}

// BookingCouponAppliedFactory ...
func BookingCouponAppliedFactory(e dto.Event) (message.Event, error) {
	var evt BookingCouponApplied
	evt.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &evt.Data); err != nil {
		return nil, err
	}

	return evt, nil
}
