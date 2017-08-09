package goinject


type IInjectable interface {
	Initialize(args ...interface{})
}

