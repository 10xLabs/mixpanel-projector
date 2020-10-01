package applier

import (
	"context"

	"github.com/10xLabs/chandler/applier"
	"github.com/10xLabs/chandler/projection"
	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/evendipiti/booking/event"
	"github.com/10xLabs/log"

	"github.com/10xLabs/mixpanel-projector/projection/booking"
)

// BookingApplier ...
type BookingApplier struct {
	*booking.Booking
}

// BookingApplierFactory ...
func BookingApplierFactory(p projection.Projection) applier.Applier {
	b := p.(*booking.Booking)

	return &BookingApplier{b}
}

// Apply ...
func (a *BookingApplier) Apply(ctx context.Context, e message.Event) error {
	log.WithFields(log.Fields{
		"aggregateID": e.AggregateID(),
		"eventID":     e.ID(),
		"type":        e.Type(),
	}).Info("applying event")

	a.ID = e.AggregateID().String()

	switch v := e.(type) {
	case event.BookingCreated:
		return a.applyBookingCreated(ctx, v)
	}

	return applier.ErrEventNotHandled
}

func (a *BookingApplier) applyBookingCreated(ctx context.Context, e event.BookingCreated) error {
	a.SessionID = e.Data.Client.SessionID.String()
	a.IP = e.Data.Client.IP

	return nil
}
