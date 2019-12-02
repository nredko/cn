package main

import (
	"github.com/codenotary/ctrlt/pkg/cmd"
	"github.com/codenotary/ctrlt/pkg/util"
)

func main() {
	if err := cmd.NewCtrlTCmd().Execute(); err != nil {
		util.Die(err)
	}
}
