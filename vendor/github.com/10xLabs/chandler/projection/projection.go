package projection

import (
	"context"

	"github.com/google/uuid"
)

// Projection ...
type Projection interface {
	Save(context.Context) error
	Load(ctx context.Context, id uuid.UUID) error
}
