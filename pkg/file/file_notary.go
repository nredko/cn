package file

import (
	"github.com/codenotary/ctrlt/pkg/notary"
)

type FileNotary interface {
	Notarize(path string, status string) (*notary.Notarization, error)
	Authenticate(path string) (*notary.Notarization, error)
}
