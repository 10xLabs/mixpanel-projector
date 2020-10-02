package projector

import (
	"context"

	"github.com/10xLabs/comandler/message"
	"github.com/10xLabs/log"

	"github.com/10xLabs/chandler/applier"
	"github.com/10xLabs/chandler/projection"
	"github.com/10xLabs/chandler/tracer"
)

var tr = tracer.New()

// Projector ...
type Projector interface {
	Project(ctx context.Context, name string, events []message.Event) error
}

// CreateProjectionFunction ...
type CreateProjectionFunction func(t string) (projection.Projection, error)

// CreateApplierFunction ...
type CreateApplierFunction func(t string, p projection.Projection) (applier.Applier, error)

type projector struct {
	createProjection CreateProjectionFunction
	createApplier    CreateApplierFunction
	skipLoadEvents   map[string]bool
}

// New ...
func New(cp CreateProjectionFunction, ca CreateApplierFunction, skipLoadEvents ...string) Projector {
	se := map[string]bool{}
	for _, e := range skipLoadEvents {
		se[e] = true
	}

	return &projector{
		createProjection: cp,
		createApplier:    ca,
		skipLoadEvents:   se,
	}
}

// Register ...
func Register(name string, pf func() projection.Projection, af func(projection.Projection) applier.Applier) {
	projection.RegisterFactory(name, pf)
	applier.RegisterFactory(name, af)
}

// Project ...
func (pr *projector) Project(ctx context.Context, name string, events []message.Event) (err error) {
	ts := tr.Trace(ctx, "projector.Project")
	defer ts.Close(err)

	if len(events) == 0 {
		log.WithFields(log.Fields{
			"name": name,
		}).Info("no events to project")

		return nil
	}

	p, err := pr.createProjection(name)
	if err != nil {
		log.WithFields(log.Fields{
			"name": name,
		}).WithError(err).Error("createProjection error")

		return Error{nil, ErrCodeCreateError, err.Error()}
	}

	a, err := pr.createApplier(name, p)
	if err != nil {
		log.WithFields(log.Fields{
			"name": name,
		}).WithError(err).Error("createApplier error")

		return Error{nil, ErrCodeCreateError, err.Error()}
	}

	fe := events[0]
	ts.AddAnnotation("aggregateID", fe.AggregateID().String())
	if fe.AggregateVersion() > 1 && !pr.skipLoadEvents[fe.Type()] {
		if err := p.Load(ctx, fe.AggregateID()); err != nil {
			log.WithFields(log.Fields{
				"aggregateID": fe.AggregateID(),
				"name":        name,
			}).WithError(err).Error("projection load error")

			return Error{fe, ErrCodeLoadError, err.Error()}
		}
	}

	for _, e := range events {
		ts.AddAnnotation(e.Type(), true)
		if err := a.Apply(ctx, e); err != nil {
			if err != applier.ErrEventNotHandled {
				log.WithFields(log.Fields{
					"aggregateID": e.AggregateID(), "event": e,
				}).WithError(err).Error("apply event error")

				return Error{e, ErrCodeApplyError, err.Error()}
			}
			log.WithFields(log.Fields{
				"aggregateID": e.AggregateID(),
				"event":       e,
			}).Info("event not handled")
		}
	}

	err = p.Save(ctx)
	if err != nil {
		return Error{nil, ErrCodeSaveError, err.Error()}
	}

	return nil
}
