package notary

import (
	"fmt"

	"github.com/codenotary/objects/pkg/object"

	"github.com/codenotary/cn/pkg/constants"
)

type StoreMeta map[string]interface{}

type Notarization struct {
	Status    string
	Object    *object.Object
	StoreMeta StoreMeta
}

var UnknownNotarization = &Notarization{
	Status:    constants.Unknown,
	Object:    nil,
	StoreMeta: nil,
}

type storedNotarization struct {
	Status string
	Object *object.Object
}

func (n Notarization) String() string {
	return fmt.Sprintf("Status:%s Object:%+v StoreMeta:%+v",
		n.Status, n.Object, n.StoreMeta)
}
