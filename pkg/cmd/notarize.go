package cmd

import (
	"os"

	"github.com/spf13/cobra"

	. "github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/printer"
	"github.com/codenotary/ctrlt/pkg/util"
)

func NewNotarizeCmd(output *string) *cobra.Command {
	return &cobra.Command{
		Use:     "notarize",
		Example: "cn notarize file://document.txt, cn notarize docker://alpine",
		Aliases: []string{"n"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			result, err := util.NotarizeSchema(args[0], Notarized)
			if err != nil {
				util.Die("notarization failed:", err)
			}
			if err = printer.Print(*output, os.Stdout, result); err != nil {
				util.Die("printing failed", err)
			}
		},
	}
}
