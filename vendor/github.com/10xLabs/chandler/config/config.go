package config

import (
	"reflect"
	"strings"
)

// Constants ...
const (
	ReplaySuffix           = "_replay"
	ProductionEnvironment  = "production"
	DevelopmentEnvironment = "development"
)

// SetReplayMode ...
func SetReplayMode(activate bool, collections interface{}) {
	values := reflect.ValueOf(collections)

	if activate {
		renameCollections(values.Elem(), ReplaySuffix)
	} else {
		renameCollections(values.Elem(), "")
	}
}

// renameCollections ...
func renameCollections(value reflect.Value, suffix string) {
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		c := strings.Replace(field.String(), ReplaySuffix, "", -1)
		field.SetString(c + suffix)
	}
}
