package parser

import (
	"context"
	"encoding/json"

	"github.com/10xLabs/chandler/awsevent"
	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/10xLabs/log"
)

// Parser ...
type Parser interface {
	Parse(context.Context, []byte) (message.Event, error)
	ParseEventRecords(ctx context.Context, records []*awsevent.EventRecord) ([]message.Event, error)
}

// NewParser ...
func NewParser() Parser {
	return &parser{}
}

// parser ...
type parser struct{}

// Parse ...
func (p *parser) Parse(ctx context.Context, data []byte) (message.Event, error) {
	var d dto.Event
	if err := json.Unmarshal(data, &d); err != nil {
		log.WithFields(log.Fields{
			"data": string(data),
		}).WithError(err).Error("the event could not be unmarshalled")

		return nil, err
	}

	e, err := message.CreateEvent(d)
	if err != nil {
		log.WithFields(log.Fields{
			"dtoEvent": d,
			"data":     string(data),
		}).WithError(err).Error("error creating event")

		return nil, err
	}

	return e, nil
}

// ParseEventRecords ...
func (p *parser) ParseEventRecords(ctx context.Context, records []*awsevent.EventRecord) ([]message.Event, error) {
	events := []message.Event{}
	for _, r := range records {
		e, err := p.Parse(ctx, r.Data())
		if err != nil {
			if err == message.ErrUnregisteredEventFactory {
				continue
			}
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}
