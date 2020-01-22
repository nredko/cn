package main

import (
	"os"

	"github.com/codenotary/immudb/pkg/client"
	"github.com/codenotary/objects/pkg/extractor"
	"github.com/codenotary/objects/pkg/extractor/file"
	"github.com/codenotary/objects/pkg/extractor/git"

	"github.com/codenotary/logger/pkg/logger"

	. "github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/docker"
	"github.com/codenotary/ctrlt/pkg/notary"
	"github.com/codenotary/ctrlt/pkg/printer"
)

var _ = (func() interface{} {
	extractor.Register(file.Scheme, file.Extract)
	extractor.Register(git.Scheme, git.Extract)
	di.RegisterOrPanic(
		di.Entry{
			Name: Logger,
			Maker: func() (interface{}, error) {
				return logger.NewSimpleLogger("ctrl-t", os.Stderr), nil
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
			Name: DockerClient,
			Maker: func() (interface{}, error) {
				return docker.NewNativeClient()
			},
		},
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
