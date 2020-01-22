package constants

import (
	"errors"
)

var (
	ErrNoSuchImage   = errors.New("no such image")
	ErrNoSuchPrinter = errors.New("no such registered printer")
)
