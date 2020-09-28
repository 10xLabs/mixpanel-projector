package event

import (
	"encoding/json"
	"fmt"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingPaidOnSite ...
type BookingPaidOnSite struct {
	message.BaseEvent
	Data BookingPaidOnSiteData
}

// BookingPaidOnSiteData ...
type BookingPaidOnSiteData struct {
	ExternalPayment BookingPaidOnSitePaymentData `json:"externalPayment"`
}

// BookingPaidOnSitePaymentData ...
type BookingPaidOnSitePaymentData struct {
	ID        uuid.UUID `json:"id"`
	Processor *string   `json:"processor"`
	Method    string    `json:"method"`
	ChargeID  *string   `json:"chargeID"`
}

// BookingPaidOnSiteFactory ...
func BookingPaidOnSiteFactory(e dto.Event) (message.Event, error) {
	var b BookingPaidOnSite
	b.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &b.Data); err != nil {
		return nil, err
	}
	fmt.Printf("unmarshaled event: %+v", b.Data)
	return b, nil
}
