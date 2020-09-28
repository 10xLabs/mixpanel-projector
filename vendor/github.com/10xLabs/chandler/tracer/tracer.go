package tracer

import (
	"context"

	"github.com/aws/aws-xray-sdk-go/xray"
)

// Tracer ...
type Tracer interface {
	Trace(ctx context.Context, name string) Segment
}

// T ...
type T struct {
}

// Segment ...
type Segment interface {
	Close(err error)
	AddMetadata(key string, value interface{})
	AddAnnotation(key string, value interface{})
}

// S ...
type S struct {
	segment *xray.Segment
}

// New ...
func New() Tracer {
	return &T{}
}

// Trace ...
func (*T) Trace(ctx context.Context, name string) Segment {
	_, ss := xray.BeginSubsegment(ctx, name)
	return &S{ss}
}

// Close ...
func (s *S) Close(err error) {
	if s.segment != nil {
		s.segment.Close(err)
	}
}

// AddMetadata ...
func (s *S) AddMetadata(key string, value interface{}) {
	s.segment.AddMetadata(key, value)
}

// AddAnnotation ...
func (s *S) AddAnnotation(key string, value interface{}) {
	s.segment.AddAnnotation(key, value)
}
