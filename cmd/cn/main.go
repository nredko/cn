package main

import (
	"github.com/codenotary/cn/pkg/cmd"
	"github.com/codenotary/cn/pkg/util"
)

func main() {
	if err := cmd.NewCnCommand().Execute(); err != nil {
		util.Die(err)
	}
}
