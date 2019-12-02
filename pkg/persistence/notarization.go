package persistence

import (
	"fmt"

	"github.com/codenotary/ctrlt/pkg/constants"
)

type Notarization struct {
	Hash   string
	Status string
	Index  uint64
}

var UnknownNotarization = &Notarization{
	Hash:   "",
	Status: constants.Unknown,
	Index:  0,
}

func (n Notarization) String() string {
	return fmt.Sprintf("Hash:%s Status:%s Index:%d",
		n.Hash, n.Status, n.Index)
}
