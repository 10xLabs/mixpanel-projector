package dto

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// DTO represents a data transfer object.
type DTO interface{}

// Message ...
type Message struct {
	ID            uuid.UUID `json:"id"`
	CorrelationID uuid.UUID `json:"correlationID"`
	Type          string    `json:"type"`
}

// Command ...
type Command struct {
	Message
	Data        json.RawMessage `json:"data"`
	AccessToken string          `json:"accessToken"`
}

// Event ...
type Event struct {
	Message
	AggregateID      uuid.UUID       `json:"aggregateID"`
	AggregateName    string          `json:"aggregateName"`
	AggregateVersion int16           `json:"aggregateVersion"`
	CreatedAt        time.Time       `json:"createdAt"`
	Data             json.RawMessage `json:"data"`
	Metadata         Metadata        `json:"metadata"`
}

// Metadata ...
type Metadata struct {
	UserID uuid.UUID `json:"userID"`
}
