package main

import (
	"os"

	"github.com/codenotary/immudb/pkg/client"

	"github.com/codenotary/logger/pkg/logger"

	"github.com/codenotary/di/pkg/di"

	"github.com/codenotary/ctrlt/pkg/api"
	. "github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/docker"
	"github.com/codenotary/ctrlt/pkg/notary"
	"github.com/codenotary/ctrlt/pkg/ui"
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
			Name: ApiServer,
			Maker: func() (interface{}, error) {
				return api.NewApiServer()
			},
		},
		di.Entry{
			Name: UiServer,
			Maker: func() (interface{}, error) {
				return ui.NewUiServer()
			},
		})
	return nil
})()
