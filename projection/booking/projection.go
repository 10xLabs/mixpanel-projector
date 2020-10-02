package booking

import (
	"context"

	"github.com/10xLabs/chandler/projection"
	"github.com/google/uuid"
)

var (
	repo Repository
)

// Amount ...
type Amount int

// Booking ...
type Booking struct {
	ID        string `json:"id"`
	IP        string `json:"ip"`
	SessionID string `json:"sessionID"`
}

// SetRepo ...
func SetRepo(r Repository) {
	repo = r
}

// Find ...
func Find(ctx context.Context, id uuid.UUID) (*Booking, error) {
	c := &Booking{}
	err := c.Load(ctx, id)

	return c, err
}

// Load ...
func (b *Booking) Load(ctx context.Context, id uuid.UUID) error {
	b.ID = id.String()
	if err := repo.Load(ctx, id.String(), b); err != nil {
		return err
	}

	return nil
}

// Save ...
func (b *Booking) Save(ctx context.Context) error {
	return repo.Replace(ctx, b.ID, b)
}

// Factory ...
func Factory() projection.Projection {
	return &Booking{}
}
