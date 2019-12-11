package container

import (
	"fmt"

	"github.com/codenotary/ctrlt/pkg/docker"
	"github.com/codenotary/ctrlt/pkg/notary"
)

type NotarizedImage struct {
	Image        docker.Image
	Notarization notary.Notarization
}

func (n NotarizedImage) String() string {
	return fmt.Sprintf("Name:%s Notarization:%v", n.Image.Name, n.Notarization)
}
