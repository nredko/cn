package di

type Registry interface {
	Register(entries ...Entry) error
	RegisterOrPanic(entries ...Entry)
	Lookup(name string) (interface{}, error)
	LookupOrPanic(name string) interface{}
	Initialize() error
	Terminate() error
}

var globalRegistry = NewSimpleRegistry()

func Register(entries ...Entry) error {
	return globalRegistry.Register(entries...)
}

func RegisterOrPanic(entries ...Entry) {
	globalRegistry.RegisterOrPanic(entries...)
}

func Lookup(name string) (interface{}, error) {
	return globalRegistry.Lookup(name)
}

func LookupOrPanic(name string) interface{} {
	return globalRegistry.LookupOrPanic(name)
}

func Initialize() error {
	return globalRegistry.Initialize()
}

func Terminate() error {
	return globalRegistry.Terminate()
}
