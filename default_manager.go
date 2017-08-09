package goinject

import (
	"reflect"
	"fmt"
	"errors"
)


func NewDefaultDiManager(definitions ...Definition) defaultManager {
	manager := defaultManager{}

	manager.store = make(map[DiKey]IInjectable)
	manager.table = make(map[DiKey]reflect.Type)

	for _, definition := range definitions{
		manager.register(definition)
	}

	return manager
}

type defaultManager struct {
	table map[DiKey]reflect.Type
	store map[DiKey]IInjectable
}


func (di *defaultManager) register(definitions ...Definition){
	for _, definition := range definitions{
		di.table[definition.Key] = definition.Type
	}
}

func (di defaultManager) Resolve(key DiKey, args ...interface{}) (IInjectable, error) {
	if obj, ok := di.store[key]; ok{
		return obj, nil
	}
	if t, ok := di.table[key]; ok{
		obj := reflect.New(t).Interface().(IInjectable)
		obj.Initialize(args)
		di.store[key] = obj
		return obj, nil
	}
	return nil, errors.New(fmt.Sprintf("Cannot resolve an unmapped type: %d", key))
}