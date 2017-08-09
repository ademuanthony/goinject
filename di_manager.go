package goinject

import (
	"reflect"
)

type IManager interface {
	// Resolve will return the stored instance of this type if it has already being initialized
	// else it will create a new instance
	Resolve(key DiKey, args ...interface{}) (IInjectable, error)
}

type Definition struct {
	Key 	DiKey
	Type 	reflect.Type
}