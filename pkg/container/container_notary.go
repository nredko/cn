package container

import (
	"github.com/codenotary/ctrlt/pkg/notary"
)

type ContainerNotary interface {
	ListNotarizedImages(query string) ([]NotarizedImage, error)
	Notarize(hash string, status string) (*notary.Notarization, error)
	NotarizeImageWithName(name string, status string) (*notary.Notarization, error)
	GetNotarizationForHash(hash string) (*notary.Notarization, error)
	GetNotarizationHistoryForHash(hash string) ([]*notary.Notarization, error)
	GetFirstNotarizationMatchingName(name string) (*notary.Notarization, error)
}
