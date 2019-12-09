package cmd

import (
	"os"

	"github.com/spf13/cobra"

	. "github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/notarization"
	"github.com/codenotary/ctrlt/pkg/printer"
	"github.com/codenotary/ctrlt/pkg/util"
)

func NewUntrustCmd(output *string) *cobra.Command {
	return &cobra.Command{
		Use:     "untrust",
		Aliases: []string{"u"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			notary := di.LookupOrPanic(Notary).(notarization.ContainerNotary)
			result, err := notary.NotarizeImageWithName(args[0], Untrusted)
			if err != nil {
				util.Die("notarization failed:", err)
			}
			if err = printer.Print(*output, os.Stdout, result); err != nil {
				util.Die("printing failed", err)
			}
		},
	}
}
