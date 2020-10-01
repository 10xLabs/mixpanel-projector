package event

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/google/uuid"
)

// BookingCreated ...
type BookingCreated struct {
	message.BaseEvent
	Data BookingCreatedData
}

// BookingCreatedData ...
type BookingCreatedData struct {
	Domain string                   `json:"domain"`
	Client BookingCreatedClientData `json:"client"`
}

// BookingCreatedClientData ...
type BookingCreatedClientData struct {
	SessionID               uuid.UUID `json:"sessionID"`
	IP                      string    `json:"ip"`
	DeviceFingerprint       *string   `json:"deviceFingerprint"`
	UserAgent               *string   `json:"userAgent"`
	OS                      *string   `json:"os"`
	Source                  string    `json:"source"`
	GoogleAnalyticsClientID *string   `json:"googleAnalyticsClientID"`
}

// BookingCreatedFactory ...
func BookingCreatedFactory(e dto.Event) (message.Event, error) {
	var rc BookingCreated
	rc.BaseEvent = message.NewBaseEvent(e)
	if err := json.Unmarshal(e.Data, &rc.Data); err != nil {
		return nil, err
	}

	return rc, nil
}
