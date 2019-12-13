package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/cobra"

	. "github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/container"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/file"
	"github.com/codenotary/ctrlt/pkg/notary"
	"github.com/codenotary/ctrlt/pkg/printer"
	"github.com/codenotary/ctrlt/pkg/util"
)

func NewNotarizeCmd(output *string) *cobra.Command {
	return &cobra.Command{
		Use:     "notarize",
		Aliases: []string{"n"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			result, err := notarize(args[0])
			if err != nil {
				util.Die("notarization failed:", err)
			}
			if err = printer.Print(*output, os.Stdout, result); err != nil {
				util.Die("printing failed", err)
			}
		},
	}
}

func notarize(arg string) (*notary.Notarization, error) {
	if strings.HasPrefix(arg, "docker://") {
		n := di.LookupOrPanic(ContainerNotary).(container.ContainerNotary)
		imageName := strings.ReplaceAll(arg, "docker://", "")
		return n.NotarizeImageWithName(imageName, Notarized)
	} else if strings.HasPrefix(arg, "file://") {
		n := di.LookupOrPanic(FileNotary).(file.FileNotary)
		path := strings.ReplaceAll(arg, "file://", "")
		return n.Notarize(path, Notarized)
	} else {
		return nil, errors.New("unrecognized notarization schema")
	}
}
