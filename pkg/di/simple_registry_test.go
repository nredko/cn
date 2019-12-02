package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type someStruct struct{}

var startCalled = false
var stopCalled = false

func (t *someStruct) Start() error {
	startCalled = true
	return nil
}

func (t *someStruct) Stop() error {
	stopCalled = true
	return nil
}

func TestRegisterAndLookup(t *testing.T) {
	registry := NewSimpleRegistry().(*simpleRegistry)
	assert.NoError(t, registry.Register(Entry{
		Name: "test",
		Maker: func() (interface{}, error) {
			return "instance", nil
		},
	}))
	instance, err := registry.Lookup("test")
	assert.NoError(t, err)
	assert.Equal(t, "instance", instance)
	assert.NotPanics(t, func() {
		_ = registry.LookupOrPanic("test")
	})
}

func TestRegisterAndLookUpMiss(t *testing.T) {
	registry := NewSimpleRegistry().(*simpleRegistry)
	assert.NoError(t, registry.Register(Entry{
		Name: "test",
		Maker: func() (interface{}, error) {
			return "instance", nil
		},
	}))
	instance, err := registry.Lookup("nope")
	assert.Error(t, err)
	assert.Nil(t, instance)
	assert.Panics(t, func() {
		_ = registry.LookupOrPanic("nope")
	})
}

func TestInitializeInstances(t *testing.T) {
	registry := NewSimpleRegistry().(*simpleRegistry)
	assert.NoError(t, registry.Register(Entry{
		Name: "someStruct",
		Maker: func() (interface{}, error) {
			return &someStruct{}, nil
		},
	}))
	assert.NoError(t, registry.Initialize())
	assert.True(t, startCalled)
}

func TestTerminateInstances(t *testing.T) {
	registry := NewSimpleRegistry().(*simpleRegistry)
	assert.NoError(t, registry.Register(Entry{
		Name: "someStruct",
		Maker: func() (interface{}, error) {
			return &someStruct{}, nil
		},
	}))
	assert.NoError(t, registry.Terminate())
	assert.True(t, stopCalled)
}
