package applier

import "errors"

// Errors ...
var (
	ErrUnregisteredApplierFactory = errors.New("unregistered applier")
	ErrEventNotHandled            = errors.New("event not handled")
)
