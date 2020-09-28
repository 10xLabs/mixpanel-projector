package message

import (
	"github.com/google/uuid"
)

// Message ...
type Message interface {
	ID() uuid.UUID
	CorrelationID() uuid.UUID
	Type() string
}
