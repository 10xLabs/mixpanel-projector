package dto

import (
	"errors"
)

// Parser errors
var (
	ErrSintax              = errors.New("sixtax error")
	ErrInvalidMessage      = errors.New("message can not be converted to a DTO")
	ErrEmptyCommandType    = errors.New("empty command type")
	ErrEmptyKinesisRecords = errors.New("empty kinesis records")
)
