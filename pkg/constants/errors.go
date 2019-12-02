package constants

import (
	"errors"
)

var (
	ErrNoSuchImage           = errors.New("no such image")
	ErrNameAlreadyRegistered = errors.New("name already registered")
	ErrNoSuchEntry           = errors.New("no such registered instance")
	ErrNoSuchPrinter         = errors.New("no such registered printer")
)
