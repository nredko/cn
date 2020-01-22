package printer

import (
	"io"

	"github.com/codenotary/di/pkg/di"

	"github.com/codenotary/ctrlt/pkg/constants"
)

type Printer interface {
	Print(io.Writer, interface{}) error
}

func Print(format string, writer io.Writer, data interface{}) error {
	var instanceName string
	switch format {
	case "text":
		instanceName = constants.TextPrinter
	case "json":
		instanceName = constants.JsonPrinter
	case "yaml":
		instanceName = constants.YamlPrinter
	default:
		return constants.ErrNoSuchPrinter
	}
	instance, err := di.Lookup(instanceName)
	if err == di.ErrNoSuchEntry {
		return constants.ErrNoSuchPrinter
	} else if err != nil {
		return err
	}
	return instance.(Printer).Print(writer, data)
}
