package dto

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// EventBuilder ...
type EventBuilder interface {
	SetCorrelationID(uuid.UUID) EventBuilder
	SetType(string) EventBuilder
	SetAggregateID(uuid.UUID) EventBuilder
	SetAggregateName(string) EventBuilder
	SetAggregateVersion(int16) EventBuilder
	SetData(json.RawMessage) EventBuilder
	SetMetadata(Metadata) EventBuilder
	Build() (*Event, error)
}

type eventBuilder struct {
	event Event
}

// NewEventBuilder ...
func NewEventBuilder() EventBuilder {
	return &eventBuilder{}
}

func (e *eventBuilder) SetCorrelationID(correlationID uuid.UUID) EventBuilder {
	e.event.CorrelationID = correlationID
	return e
}

func (e *eventBuilder) SetType(etype string) EventBuilder {
	e.event.Type = etype
	return e
}

func (e *eventBuilder) SetAggregateID(aggregateID uuid.UUID) EventBuilder {
	e.event.AggregateID = aggregateID
	return e
}

func (e *eventBuilder) SetAggregateName(aggregateName string) EventBuilder {
	e.event.AggregateName = aggregateName
	return e
}

func (e *eventBuilder) SetAggregateVersion(aggregateVersion int16) EventBuilder {
	e.event.AggregateVersion = aggregateVersion
	return e
}

func (e *eventBuilder) SetData(data json.RawMessage) EventBuilder {
	e.event.Data = data
	return e
}

func (e *eventBuilder) SetMetadata(m Metadata) EventBuilder {
	e.event.Metadata = m
	return e
}

func (e *eventBuilder) Build() (*Event, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	e.event.Message.ID = id
	e.event.CreatedAt = time.Now().UTC()

	return &e.event, nil
}
