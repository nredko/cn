package notary

import (
	"github.com/codenotary/objects/pkg/object"
)

type Notary interface {
	Start() error
	Stop() error
	Notarize(object *object.Object, status string) (*Notarization, error)
	Authenticate(object *object.Object) (*Notarization, error)
	AuthenticateBatch(objects []*object.Object) ([]Notarization, error)
	History(object *object.Object) ([]*Notarization, error)
}
