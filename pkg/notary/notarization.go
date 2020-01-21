package notary

import (
	"fmt"

	"github.com/codenotary/ctrlt/pkg/constants"
)

type Meta map[string]interface{}

type StoreMeta map[string]interface{}

type Notarization struct {
	Hash      string
	Status    string
	Meta      Meta
	StoreMeta StoreMeta
}

var UnknownNotarization = &Notarization{
	Hash:      "",
	Status:    constants.Unknown,
	Meta:      nil,
	StoreMeta: nil,
}

type storedNotarization struct {
	Hash   string
	Status string
	Meta   Meta
}

func (n Notarization) String() string {
	return fmt.Sprintf("Hash:%s Status:%s Meta:%+v StoreMeta:%+v",
		n.Hash, n.Status, n.Meta, n.StoreMeta)
}
