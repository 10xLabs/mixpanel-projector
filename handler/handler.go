package handler

import (
	"context"

	capplier "github.com/10xLabs/chandler/applier"
	"github.com/10xLabs/chandler/awsevent"
	"github.com/10xLabs/chandler/parser"
	"github.com/10xLabs/chandler/projection"
	"github.com/10xLabs/chandler/projector"
	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/evendipiti/booking/event"
	"github.com/10xLabs/log"

	"github.com/10xLabs/mixpanel-projector/applier"
	"github.com/10xLabs/mixpanel-projector/config"
	"github.com/10xLabs/mixpanel-projector/projection/booking"
	"github.com/10xLabs/mixpanel-projector/projection/mixevent"
	"github.com/10xLabs/mixpanel-projector/subscriber"
)

const (
	mixeventProjection = "Mixevent"
	bookingProjection  = "Booking"
)

// Event ...
type Event struct {
	Event      string      `json:"event"`
	Properties interface{} `json:"properties"`
}

// EventProperties ...
type EventProperties struct {
	BookingID string `json:"bookingID"`
	Token     string `json:"token"`
}

// Setup ...
func Setup() {
	subscriber.Subscribe()
	booking.SetRepo(booking.NewRepository(config.App.Store, &config.App.FileStore.Dir))

	projector.Register(mixeventProjection, mixevent.Factory, applier.MixeventApplierFactory)
	projector.Register(bookingProjection, booking.Factory, applier.BookingApplierFactory)
}

var (
	p  = projector.New(projection.Create, capplier.Create)
	pa = parser.NewParser()
)

// Handler ...
func Handler(ctx context.Context, ae awsevent.Event) error {
	for _, r := range ae.Records {
		e, err := pa.Parse(ctx, r.Data())
		if err != nil {
			log.WithError(err).Error("parse error")

			return err
		}
		if err := p.Project(ctx, mixeventProjection, []message.Event{e}); err != nil {
			err2 := err.(projector.Error)
			log.WithFields(log.Fields{
				"aggregateID": e.AggregateID(),
				"err2":        err2,
			}).WithError(err).Error("mixevent project error")

			return err
		}

		if e.Type() != event.TypeBookingCreated {
			continue
		}

		if err := p.Project(ctx, bookingProjection, []message.Event{e}); err != nil {
			err2 := err.(projector.Error)
			log.WithFields(log.Fields{
				"aggregateID": e.AggregateID(),
				"err2":        err2,
			}).WithError(err).Error("booking project error")

			return err
		}
	}

	return nil
}
