package cmd

import (
	"os"

	"github.com/spf13/cobra"

	. "github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/container"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/printer"
	"github.com/codenotary/ctrlt/pkg/util"
)

func NewVerifyCmd(output *string) *cobra.Command {
	return &cobra.Command{
		Use:     "verify",
		Aliases: []string{"v"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			notary := di.LookupOrPanic(ContainerNotary).(container.ContainerNotary)
			result, err := notary.GetFirstNotarizationMatchingName(args[0])
			if err != nil {
				util.Die("verification failed:", err)
			}
			if err = printer.Print(*output, os.Stdout, *result); err != nil {
				util.Die("printing failed", err)
			}
		},
	}
}
