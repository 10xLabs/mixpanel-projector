package applier

import (
	"context"

	"github.com/10xLabs/chandler/applier"
	"github.com/10xLabs/chandler/projection"
	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/comandler/message/dto"
	"github.com/10xLabs/evendipiti/booking/event"
	"github.com/10xLabs/log"

	"github.com/10xLabs/mixpanel-projector/projection/mixevent"
)

// MixeventApplier ...
type MixeventApplier struct {
	*mixevent.Mixevent
}

// MixeventApplierFactory ...
func MixeventApplierFactory(p projection.Projection) applier.Applier {
	u := p.(*mixevent.Mixevent)

	return &MixeventApplier{u}
}

// Apply ...
func (a *MixeventApplier) Apply(ctx context.Context, e message.Event) error {
	log.WithFields(log.Fields{
		"aggregateID": e.AggregateID(),
		"eventID":     e.ID(),
		"type":        e.Type(),
	}).Info("applying event")

	a.ID = e.ID().String()
	a.Version = e.AggregateVersion()
	a.Event = mixevent.Event{
		Name: e.Type(),
		Properties: mixevent.EventProperties{
			InsertID: e.ID().String(),
			Time:     e.CreatedAt().UnixNano(),
			Name:     e.Type(),
		},
	}
	switch v := e.(type) {
	case event.BookingCreated:
		return a.applyBookingCreated(ctx, v)
	}

	return a.applyBookingEvent(ctx, e.DTO())
}

func (a *MixeventApplier) applyBookingEvent(ctx context.Context, e dto.Event) error {
	a.Event.Properties.Data = e.Data
	bk, err := bookingFind(ctx, e.AggregateID)
	if err != nil {
		return err
	}
	a.Event.Properties.DistinctID = bk.SessionID
	a.Event.Properties.IP = bk.IP

	return nil
}

func (a *MixeventApplier) applyBookingCreated(ctx context.Context, e event.BookingCreated) error {
	a.Event.Properties.Data = e.Data
	a.Event.Properties.IP = e.Data.Client.IP

	return nil
}
