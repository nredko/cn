package cmd

import (
	"os"

	"github.com/codenotary/objects/pkg/extractor"
	"github.com/spf13/cobra"

	. "github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/notary"
	"github.com/codenotary/ctrlt/pkg/printer"
	"github.com/codenotary/ctrlt/pkg/util"
)

func NewUntrustCmd(output *string) *cobra.Command {
	return &cobra.Command{
		Use:     "untrust",
		Example: "cn untrust file://document.txt",
		Aliases: []string{"u"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			n := di.LookupOrPanic(Notary).(notary.Notary)
			o, err := extractor.Extract(args[0])
			if err != nil {
				util.Die("notarization failed:", err)
			}
			result, err := n.Notarize(o, Untrusted)
			if err != nil {
				util.Die("notarization failed:", err)
			}
			if err = printer.Print(*output, os.Stdout, result); err != nil {
				util.Die("printing failed", err)
			}
		},
	}
}
