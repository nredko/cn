package cmd

import (
	"github.com/spf13/cobra"

	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/util"
)

func NewCnCommand() *cobra.Command {
	var output string
	cmd := &cobra.Command{
		Use: "cn",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := di.Initialize(); err != nil {
				util.Die("initializing failed:", err)
			}
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if err := di.Terminate(); err != nil {
				util.Die("termination failed:", err)
			}
		},
	}
	cmd.PersistentFlags().StringVarP(&output, "output", "o", "text", "output")
	cmd.AddCommand(
		NewNotarizeCmd(&output),
		NewUntrustCmd(&output),
		NewVerifyCmd(&output),
		NewListCmd(&output))
	return cmd
}
