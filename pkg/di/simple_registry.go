package di

import (
	"os"
	"reflect"
	"time"

	"github.com/codenotary/logger/pkg/logger"

	"github.com/codenotary/ctrlt/pkg/constants"
)

type simpleRegistry struct {
	instances map[string]interface{}
	logger    logger.Logger
}

func NewSimpleRegistry() Registry {
	return &simpleRegistry{
		instances: map[string]interface{}{},
		logger:    logger.NewSimpleLogger("ctrl-t", os.Stderr),
	}
}

func (r *simpleRegistry) Register(entries ...Entry) error {
	start := time.Now()
	for _, entry := range entries {
		if r.instances[entry.Name] != nil {
			return constants.ErrNameAlreadyRegistered
		}
		if instance, err := entry.Maker(); err != nil {
			return err
		} else {
			r.logger.Debugf("registered '%s' (%s@%p)", entry.Name,
				reflect.TypeOf(instance), instance)
			r.instances[entry.Name] = instance
		}
	}
	r.logger.Debugf("registration finished in %s", time.Since(start))
	return nil
}

func (r *simpleRegistry) RegisterOrPanic(entries ...Entry) {
	if err := r.Register(entries...); err != nil {
		r.logger.Errorf("register failed: %v", err)
		panic(err)
	}
}

func (r *simpleRegistry) Lookup(name string) (interface{}, error) {
	if instance := r.instances[name]; instance == nil {
		return nil, constants.ErrNoSuchEntry
	} else {
		return instance, nil
	}
}

func (r *simpleRegistry) LookupOrPanic(name string) interface{} {
	instance, err := r.Lookup(name)
	if err != nil {
		r.logger.Errorf("lookup '%s' failed: %v", name, err)
		panic(err)
	}
	return instance
}

func (r *simpleRegistry) Initialize() error {
	globalStart := time.Now()
	for name, instance := range r.instances {
		initializingType := reflect.TypeOf((*Initializing)(nil)).Elem()
		if reflect.TypeOf(instance).Implements(initializingType) {
			start := time.Now()
			if err := instance.(Initializing).Start(); err != nil {
				return err
			}
			r.logger.Debugf("initialized '%s' (%s@%p) in %s", name,
				reflect.TypeOf(instance), instance, time.Since(start))
		}
	}
	r.logger.Debugf("initialization finished in %s", time.Since(globalStart))
	return nil
}

func (r *simpleRegistry) Terminate() error {
	globalStart := time.Now()
	for name, instance := range r.instances {
		terminatingType := reflect.TypeOf((*Terminating)(nil)).Elem()
		if reflect.TypeOf(instance).Implements(terminatingType) {
			start := time.Now()
			if err := instance.(Terminating).Stop(); err != nil {
				return err
			}
			r.logger.Debugf("terminated '%s' (%s@%p) in %s", name,
				reflect.TypeOf(instance), instance, time.Since(start))
		}
	}
	r.logger.Debugf("termination finished in %s", time.Since(globalStart))
	return nil
}
