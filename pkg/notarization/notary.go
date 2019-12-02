package notarization

import (
	"github.com/codenotary/ctrlt/pkg/persistence"
)

type Notary interface {
	ListNotarizedImages(query string) ([]NotarizedImage, error)
	Notarize(hash string, status string) (*persistence.Notarization, error)
	NotarizeImageWithName(name string, status string) (*persistence.Notarization, error)
	GetNotarizationForHash(hash string) (*persistence.Notarization, error)
	GetFirstNotarizationMatchingName(name string) (*persistence.Notarization, error)
}
