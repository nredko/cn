package printer

import (
	"encoding/json"
	"io"
)

type JsonPrinter struct{}

func NewJsonPrinter() (Printer, error) {
	return &JsonPrinter{}, nil
}

func (p *JsonPrinter) Print(writer io.Writer, data interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = writer.Write(payload)
	return err
}
