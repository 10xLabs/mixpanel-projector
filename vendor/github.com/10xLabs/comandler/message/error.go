package message

import "errors"

// Message errors
var (
	ErrSintax                     = errors.New("sixtax error")
	ErrUnregisteredCommandFactory = errors.New("unregistered command factory")
	ErrUnregisteredEventFactory   = errors.New("unregistered event factory")
)
