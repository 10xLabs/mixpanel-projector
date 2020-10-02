package message

import (
	"fmt"

	"github.com/10xLabs/comandler/message/dto"
	"github.com/10xLabs/log"
)

type commandFactory func(dto.Command) (Command, error)

var commandFactoryRegister = map[string]commandFactory{}

// RegisterCommandFactory ...
func RegisterCommandFactory(commandType string, f commandFactory) {
	if f == nil {
		log.Panicf("command factory for %s is nil", commandType)
	}

	if _, ok := commandFactoryRegister[commandType]; ok {
		fmt.Printf("command factory for %s will be overwritten", commandType)
	}

	commandFactoryRegister[commandType] = f
}

// CreateCommand ...
func CreateCommand(c dto.Command) (Command, error) {
	factory := commandFactoryRegister[c.Type]

	if factory == nil {
		return nil, ErrUnregisteredCommandFactory
	}

	return factory(c)
}
