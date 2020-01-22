package main

import (
	"os"

	"github.com/codenotary/immudb/pkg/client"
	"github.com/codenotary/logger/pkg/logger"

	"github.com/codenotary/di/pkg/di"

	. "github.com/codenotary/cn/pkg/constants"
	"github.com/codenotary/cn/pkg/notary"
	"github.com/codenotary/cn/pkg/printer"
)

var _ = (func() interface{} {
	di.RegisterOrPanic(
		di.Entry{
			Name: Logger,
			Maker: func() (interface{}, error) {
				return logger.NewSimpleLogger("cn", os.Stderr), nil
			},
		},
		di.Entry{
			Name: ImmuClient,
			Maker: func() (interface{}, error) {
				return client.DefaultClient().WithOptions(
					client.DefaultOptions().
						WithDialRetries(0).
						WithHealthCheckRetries(0).
						FromEnvironment()), nil
			},
		},
		di.Entry{
			Name: Notary,
			Maker: func() (interface{}, error) {
				return notary.NewImmuNotary()
			}},
		di.Entry{
			Name: TextPrinter,
			Maker: func() (interface{}, error) {
				return printer.NewTextPrinter()
			},
		},
		di.Entry{
			Name: JsonPrinter,
			Maker: func() (interface{}, error) {
				return printer.NewJsonPrinter()
			},
		},
		di.Entry{
			Name: YamlPrinter,
			Maker: func() (interface{}, error) {
				return printer.NewYamlPrinter()
			},
		})
	return nil
})()
