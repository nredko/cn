package printer

import (
	"fmt"
	"io"
)

type TextPrinter struct{}

func NewTextPrinter() (Printer, error) {
	return &TextPrinter{}, nil
}

func (p *TextPrinter) Print(writer io.Writer, data interface{}) error {
	_, err := fmt.Fprint(writer, fmt.Sprintln(data))
	return err
}
