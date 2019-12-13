package util

import (
	"strings"

	. "github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/container"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/file"
	"github.com/codenotary/ctrlt/pkg/notary"
)

func NotarizeSchema(arg string, status string) (*notary.Notarization, error) {
	if strings.HasPrefix(arg, "docker://") {
		n := di.LookupOrPanic(ContainerNotary).(container.ContainerNotary)
		imageName := strings.ReplaceAll(arg, "docker://", "")
		return n.NotarizeImageWithName(imageName, status)
	} else if strings.HasPrefix(arg, "file://") {
		n := di.LookupOrPanic(FileNotary).(file.FileNotary)
		path := strings.ReplaceAll(arg, "file://", "")
		return n.Notarize(path, status)
	} else {
		n := di.LookupOrPanic(FileNotary).(file.FileNotary)
		return n.Notarize(arg, status)
	}
}

func VerifySchema(arg string) (*notary.Notarization, error) {
	if strings.HasPrefix(arg, "docker://") {
		n := di.LookupOrPanic(ContainerNotary).(container.ContainerNotary)
		imageName := strings.ReplaceAll(arg, "docker://", "")
		return n.GetFirstNotarizationMatchingName(imageName)
	} else if strings.HasPrefix(arg, "file://") {
		n := di.LookupOrPanic(FileNotary).(file.FileNotary)
		path := strings.ReplaceAll(arg, "file://", "")
		return n.Authenticate(path)
	} else {
		n := di.LookupOrPanic(FileNotary).(file.FileNotary)
		return n.Authenticate(arg)
	}
}
