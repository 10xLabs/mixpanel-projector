package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingCreatedOnSite ...
type BookingCreatedOnSite struct {
	message.BaseEvent
	Data BookingCreatedOnSiteData
}

// BookingCreatedOnSiteData ...
type BookingCreatedOnSiteData struct {
	Domain string    `json:"domain"`
	SiteID uuid.UUID `json:"siteID"`
}

// BookingCreatedOnSiteFactory ...
func BookingCreatedOnSiteFactory(e dto.Event) (message.Event, error) {
	var bc BookingCreatedOnSite
	bc.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &bc.Data); err != nil {
		return nil, err
	}

	return bc, nil
}
