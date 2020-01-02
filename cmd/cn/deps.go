package main

import (
	"os"

	"github.com/codenotary/immudb/pkg/client"

	. "github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/container"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/docker"
	"github.com/codenotary/ctrlt/pkg/file"
	"github.com/codenotary/ctrlt/pkg/logger"
	"github.com/codenotary/ctrlt/pkg/notary"
	"github.com/codenotary/ctrlt/pkg/printer"
)

var _ = (func() interface{} {
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
			Name: ContainerNotary,
			Maker: func() (interface{}, error) {
				return container.NewDockerNotary()
			},
		},
		di.Entry{
			Name: FileNotary,
			Maker: func() (interface{}, error) {
				return file.NewLocalFileNotary()
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
