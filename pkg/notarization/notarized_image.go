package notarization

import (
	"fmt"

	"github.com/codenotary/ctrlt/pkg/docker"
	"github.com/codenotary/ctrlt/pkg/persistence"
)

type NotarizedImage struct {
	Image        docker.Image
	Notarization persistence.Notarization
}

func (n NotarizedImage) String() string {
	return fmt.Sprintf("Name:%s Notarization:%v", n.Image.Name, n.Notarization)
}
