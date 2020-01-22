package cmd

import (
	"os"

	"github.com/codenotary/objects/pkg/extractor"
	"github.com/spf13/cobra"

	"github.com/codenotary/di/pkg/di"

	"github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/notary"
	"github.com/codenotary/ctrlt/pkg/printer"
	"github.com/codenotary/ctrlt/pkg/util"
)

func NewHistoryCmd(output *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "history",
		Aliases: []string{"h"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			n := di.LookupOrPanic(constants.Notary).(notary.Notary)
			o, err := extractor.Extract(args[0])
			if err != nil {
				util.Die("history retrieval failed:", err)
			}
			history, err := n.History(o)
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
