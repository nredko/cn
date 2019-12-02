package util

import (
	"fmt"
	"os"
)

func Die(args ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}
