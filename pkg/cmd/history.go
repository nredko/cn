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

func NewHistoryCmd(output *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "history",
		Aliases: []string{"h"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			notary := di.LookupOrPanic(Notary).(notarization.ContainerNotary)
			notarizedImage, err := notary.GetFirstNotarizationMatchingName(args[0])
			if err != nil {
				util.Die("notarization failed:", err)
			}
			history, err := notary.GetNotarizationHistoryForHash(notarizedImage.Hash)
			if err != nil {
				util.Die("history retrieval failed:", err)
			}
			if err = printer.Print(*output, os.Stdout, history); err != nil {
				util.Die("printing failed", err)
			}
		},
	}
	return cmd
}
