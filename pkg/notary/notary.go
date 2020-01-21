package notary

import (
	"github.com/codenotary/objects/pkg/object"
)

//go:generate mockgen -source=../notary/notary.go -destination=../mocks/mock_notary.go -package=mocks

type Notary interface {
	Start() error
	Notarize(object *object.Object, status string) (*Notarization, error)
	Authenticate(object *object.Object) (*Notarization, error)
	AuthenticateBatch(objects []*object.Object) ([]Notarization, error)
	History(object *object.Object) ([]*Notarization, error)
}
