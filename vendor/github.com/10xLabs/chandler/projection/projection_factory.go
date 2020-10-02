package projection

import (
	"github.com/10xLabs/log"
)

type factory func() Projection

var factoryRegister = map[string]factory{}

// Create ...
func Create(projectionType string) (Projection, error) {
	if f, ok := factoryRegister[projectionType]; ok {
		return f(), nil
	}

	return nil, ErrUnregisteredFactory
}

// RegisterFactory ...
func RegisterFactory(projectionType string, f factory) {
	if f == nil {
		log.Panicf("projection factory for %s is nil", projectionType)
	}

	if _, ok := factoryRegister[projectionType]; ok {
		log.Infof("projection factory for %s will be overwritten", projectionType)
	}

	factoryRegister[projectionType] = f
}
