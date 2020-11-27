package projector

import "github.com/10xLabs/comandler/message"

// ErrorCodes ...
const (
	ErrCodeLoadError   = "load_error"
	ErrCodeCreateError = "create_error"
	ErrCodeApplyError  = "apply_error"
	ErrCodeSaveError   = "save_error"
)

// Error ...
type Error struct {
	Event   message.Event
	Code    string
	Message string
}

func (e Error) Error() string {
	return e.Code
}
