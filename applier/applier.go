package applier

import (
	"context"

	"github.com/10xLabs/mixpanel-projector/projection/booking"
	"github.com/google/uuid"
)

var findBooking func(ctx context.Context, id uuid.UUID) (*booking.Booking, error) = booking.Find
