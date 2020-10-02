package message

import (
	"encoding/json"
	"time"

	"github.com/10xLabs/comandler/message/dto"

	"github.com/google/uuid"
)

// Event ...
type Event interface {
	Message
	AggregateID() uuid.UUID
	AggregateName() string
	AggregateVersion() int16
	ConflictWith(Event) bool
	CreatedAt() time.Time
	RawData() json.RawMessage
	DTO() dto.Event
}

// BaseEvent ...
type BaseEvent struct {
	event dto.Event
}

// NewBaseEvent ...
func NewBaseEvent(e dto.Event) BaseEvent {
	return BaseEvent{e}
}

// ID ...
func (e BaseEvent) ID() uuid.UUID {
	return e.event.ID
}

// CorrelationID ...
func (e BaseEvent) CorrelationID() uuid.UUID {
	return e.event.CorrelationID
}

// Type ...
func (e BaseEvent) Type() string {
	return e.event.Type
}

// AggregateID ...
func (e BaseEvent) AggregateID() uuid.UUID {
	return e.event.AggregateID
}

// AggregateName ...
func (e BaseEvent) AggregateName() string {
	return e.event.AggregateName
}

// AggregateVersion ...
func (e BaseEvent) AggregateVersion() int16 {
	return e.event.AggregateVersion
}

// ConflictWith ...
func (BaseEvent) ConflictWith(e Event) bool {
	return true
}

// CreatedAt ...
func (e BaseEvent) CreatedAt() time.Time {
	return e.event.CreatedAt
}

// RawData ...
func (e BaseEvent) RawData() json.RawMessage {
	return e.event.Data
}

// DTO ...
func (e BaseEvent) DTO() dto.Event {
	return e.event
}

// UserID ...
func (e BaseEvent) UserID() uuid.UUID {
	return e.event.Metadata.UserID
}
