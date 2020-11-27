package booking

import (
	rp "github.com/10xLabs/chandler/repository"
	"github.com/10xLabs/chandler/store"
)

// Repository ...
type Repository interface {
	rp.Repository
}

type repository struct {
	*rp.Base
}

// NewRepository ...
func NewRepository(s store.Store, dir *string) Repository {
	return &repository{&rp.Base{Store: s, CollectionName: dir}}
}
