package util

import (
	"io/ioutil"
	"os"

	"github.com/codenotary/immustore/pkg/server"
)

func WithImmuServer(f func(immuServer *server.ImmuServer) error) error {
	tmpDir, err := ioutil.TempDir("", "immudb")
	if err != nil {
		return err
	}
	options := server.DefaultOptions().WithDir(tmpDir)
	immuServer := server.DefaultServer().WithOptions(options)
	go immuServer.Start()
	if err = f(immuServer); err != nil {
		return err
	}
	if err = immuServer.Stop(); err != nil {
		return err
	}
	return os.RemoveAll(immuServer.Options.Dir)
}
