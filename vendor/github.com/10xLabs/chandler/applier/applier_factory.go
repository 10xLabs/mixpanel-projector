package applier

import (
	"github.com/10xLabs/log"

	"github.com/10xLabs/chandler/projection"
)

type factory func(projection.Projection) Applier

var factoryRegister = map[string]factory{}

// Create ...
func Create(applierType string, p projection.Projection) (Applier, error) {
	if f, ok := factoryRegister[applierType]; ok {
		return f(p), nil
	}

	return nil, ErrUnregisteredApplierFactory
}

// RegisterFactory ...
func RegisterFactory(applierType string, f factory) {
	if f == nil {
		log.Panicf("applier factory for %s is nil", applierType)
	}

	if _, ok := factoryRegister[applierType]; ok {
		log.Infof("applier factory for %s will be overwritten", applierType)
	}

	factoryRegister[applierType] = f
}
