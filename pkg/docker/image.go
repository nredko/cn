package docker

import (
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
)

type Image struct {
	Name string
	Hash string
}

func fromContainer(container types.Container) Image {
	return Image{
		Name: container.Image,
		Hash: strings.ReplaceAll(container.ImageID, "sha256:", ""),
	}
}

func (i Image) String() string {
	return fmt.Sprintf("Name:%s Hash:%s", i.Name, i.Hash)
}
