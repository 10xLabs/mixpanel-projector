package dto

import (
	"encoding/json"

	"github.com/google/uuid"
)

// CommandBuilder ...
type CommandBuilder interface {
	SetAccessToken(string) CommandBuilder
	SetCorrelationID(uuid.UUID) CommandBuilder
	SetData(json.RawMessage) CommandBuilder
	SetType(string) CommandBuilder
	Build() (*Command, error)
}

type commandBuilder struct {
	command Command
}

// NewCommandBuilder ...
func NewCommandBuilder() CommandBuilder {
	return &commandBuilder{}
}

func (b *commandBuilder) SetAccessToken(a string) CommandBuilder {
	b.command.AccessToken = a
	return b
}

func (b *commandBuilder) SetCorrelationID(c uuid.UUID) CommandBuilder {
	b.command.CorrelationID = c
	return b
}

func (b *commandBuilder) SetData(d json.RawMessage) CommandBuilder {
	b.command.Data = d
	return b
}

func (b *commandBuilder) SetType(t string) CommandBuilder {
	b.command.Type = t
	return b
}

func (b *commandBuilder) Build() (*Command, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	b.command.Message.ID = id

	return &b.command, nil
}
