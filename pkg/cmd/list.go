package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/codenotary/ctrlt/pkg/printer"
	"github.com/codenotary/ctrlt/pkg/util"
)

func NewListCmd(output *string) *cobra.Command {
	var query string
	var schema string
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Example: "cn list -q alpine -s docker://",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			notarizedImages, err := util.List(schema, query)
			if err != nil {
				util.Die("listing failed:", err)
			}
			if err = printer.Print(*output, os.Stdout, notarizedImages); err != nil {
				util.Die("printing failed", err)
			}
		},
	}
	cmd.Flags().StringVarP(&query, "query", "q", "", "query")
	cmd.Flags().StringVarP(&schema, "schema", "s", "docker://", "query")
	return cmd
}
