package notary

import (
	"fmt"

	"github.com/codenotary/ctrlt/pkg/constants"
)

type Notarization struct {
	Hash      string
	Status    string
	Meta      interface{}
	StoreMeta interface{}
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
	Meta   interface{}
}

func (n Notarization) String() string {
	return fmt.Sprintf("Hash:%s Status:%s Meta:%+v StoreMeta:%+v",
		n.Hash, n.Status, n.Meta, n.StoreMeta)
}
