package applier

import (
	"context"

	"github.com/google/uuid"

	"github.com/10xLabs/mixpanel-projector/projection/booking"
)

var bookingFind func(ctx context.Context, id uuid.UUID) (*booking.Booking, error) = booking.Find
