package goinject

import (
	"testing"
	"reflect"
)

type iTestInjectable interface {
	TestMethod()
}

type testInjectable struct {

}

func (t testInjectable) TestMethod() {

}

func (t testInjectable) Initialize(args ...interface{})  {

}

func TestDefaultManager_Resolve(t *testing.T) {
	var testKey DiKey = 1

	manager := NewDefaultDiManager(Definition{Key:testKey, Type:reflect.TypeOf(testInjectable{})})

	testObj, err := manager.Resolve(testKey)

	if err != nil{
		t.Error(err)
	}

	if testObj == nil{
		t.Error("Expected an obj, got nil")
	}

	if _, ok := testObj.(iTestInjectable); !ok {
		t.Error("Expected %t, got %t", testInjectable{}, testObj)
	}

}
