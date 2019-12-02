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

func NewListCmd(output *string) *cobra.Command {
	var query string
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			notary := di.LookupOrPanic(Notary).(notarization.Notary)
			notarizedImages, err := notary.ListNotarizedImages(query)
			if err != nil {
				util.Die("listing failed:", err)
			}
			if err = printer.Print(*output, os.Stdout, notarizedImages); err != nil {
				util.Die("printing failed", err)
			}
		},
	}
	cmd.Flags().StringVarP(&query, "query", "q", "", "query")
	return cmd
}
