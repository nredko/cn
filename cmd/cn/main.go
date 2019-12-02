package main

import (
	"github.com/codenotary/ctrlt/pkg/cmd"
	"github.com/codenotary/ctrlt/pkg/util"
)

func main() {
	if err := cmd.NewCnCommand().Execute(); err != nil {
		util.Die(err)
	}
}
