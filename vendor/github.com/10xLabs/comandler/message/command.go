package message

import (
	"encoding/json"

	"github.com/10xLabs/comandler/message/dto"

	"github.com/google/uuid"
)

// Command ...
type Command interface {
	Message
	RawData() json.RawMessage
}

// BaseCommand ...
type BaseCommand struct {
	command dto.Command
}

// NewBaseCommand ...
func NewBaseCommand(c dto.Command) BaseCommand {
	return BaseCommand{c}
}

// ID ...
func (c BaseCommand) ID() uuid.UUID {
	return c.command.ID
}

// CorrelationID ...
func (c BaseCommand) CorrelationID() uuid.UUID {
	return c.command.CorrelationID
}

// Type ...
func (c BaseCommand) Type() string {
	return c.command.Type
}

// RawData ...
func (c BaseCommand) RawData() json.RawMessage {
	return c.command.Data
}
