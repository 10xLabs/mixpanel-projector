package mixevent

import (
	"context"

	"github.com/10xLabs/chandler/projection"
	"github.com/google/uuid"

	"github.com/10xLabs/mixpanel-projector/config"
	"github.com/10xLabs/mixpanel-projector/mixpanel"
)

const apiURL = "https://api.mixpanel.com/"

// Mixevent ...
type Mixevent struct {
	ID      string `bson:"_id"`
	Version int16  `bson:"version"`
	Event   Event  `bson:"event"`
}

// Event ...
type Event struct {
	Name       string          `json:"event"`
	Properties EventProperties `json:"properties"`
}

// EventProperties ...
type EventProperties struct {
	Token      string      `json:"token"`
	DistinctID string      `json:"distinct_id,omitempty"`
	Time       int64       `json:"time"`
	Data       interface{} `json:"data"`
	InsertID   string      `json:"$insert_id,omitempty"`
	Name       string      `json:"name"`
	IP         string      `json:"ip"`
}

// Factory ...
func Factory() projection.Projection {
	return &Mixevent{}
}

// Save ...
func (b *Mixevent) Save(ctx context.Context) error {
	m := mixpanel.NewMixpanel(apiURL)
	b.Event.Properties.Token = config.App.MixpanelToken

	return m.Track(b.Event)
}

// Load ...
func (b *Mixevent) Load(ctx context.Context, id uuid.UUID) error {
	*b = Mixevent{}
	return nil
}
