package cmd

import (
	"os"

	"github.com/codenotary/objects/pkg/extractor"
	"github.com/spf13/cobra"

	"github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/notary"
	"github.com/codenotary/ctrlt/pkg/printer"
	"github.com/codenotary/ctrlt/pkg/util"
)

func NewVerifyCmd(output *string) *cobra.Command {
	return &cobra.Command{
		Use:     "verify",
		Example: "cn verify file://document.txt",
		Aliases: []string{"v"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			n := di.LookupOrPanic(constants.Notary).(notary.Notary)
			o, err := extractor.Extract(args[0])
			if err != nil {
				util.Die("verification failed:", err)
			}
			result, err := n.Authenticate(o)
			if err != nil {
				util.Die("verification failed:", err)
			}
			if err = printer.Print(*output, os.Stdout, *result); err != nil {
				util.Die("printing failed", err)
			}
		},
	}
}
