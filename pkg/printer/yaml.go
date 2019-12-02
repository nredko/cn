package printer

import (
	"io"

	"gopkg.in/yaml.v2"
)

type YamlPrinter struct{}

func NewYamlPrinter() (Printer, error) {
	return &YamlPrinter{}, nil
}

func (y *YamlPrinter) Print(writer io.Writer, data interface{}) error {
	payload, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	_, err = writer.Write(payload)
	return err
}
