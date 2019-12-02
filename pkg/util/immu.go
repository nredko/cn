package util

import (
	"io/ioutil"

	"github.com/codenotary/immudb/pkg/server"
)

type Immu struct{}

func NewImmuDB() error {
	tmpDir, err := ioutil.TempDir("", "immudb")
	if err != nil {
		return err
	}
	options := server.DefaultOptions().WithDir(tmpDir)
	if err = server.DefaultServer().WithOptions(options).Start(); err != nil {
		return err
	}
	return nil
}
