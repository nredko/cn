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

func NewVerifyCmd(output *string) *cobra.Command {
	return &cobra.Command{
		Use:     "verify",
		Aliases: []string{"v"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			result, err := verify(args[0])
			if err != nil {
				util.Die("verification failed:", err)
			}
			if err = printer.Print(*output, os.Stdout, *result); err != nil {
				util.Die("printing failed", err)
			}
		},
	}
}

func verify(arg string) (*notary.Notarization, error) {
	if strings.HasPrefix(arg, "docker://") {
		n := di.LookupOrPanic(ContainerNotary).(container.ContainerNotary)
		imageName := strings.ReplaceAll(arg, "docker://", "")
		return n.GetFirstNotarizationMatchingName(imageName)
	} else if strings.HasPrefix(arg, "file://") {
		n := di.LookupOrPanic(FileNotary).(file.FileNotary)
		path := strings.ReplaceAll(arg, "file://", "")
		return n.Authenticate(path)
	} else {
		return nil, errors.New("unrecognized notarization schema")
	}
}
