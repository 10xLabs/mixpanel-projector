package applier

import (
	"context"

	"github.com/10xLabs/comandler/message"
)

// Applier ...
type Applier interface {
	Apply(context.Context, message.Event) error
}
