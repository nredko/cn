package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/logger"
)

var ctx = context.Background()
var listOptions = types.ContainerListOptions{}

type nativeClient struct {
	logger       logger.Logger
	dockerClient *client.Client
}

func NewNativeClient() (Client, error) {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	log, err := di.Lookup(constants.Logger)
	if err != nil {
		return nil, err
	}
	return &nativeClient{
		logger:       log.(logger.Logger),
		dockerClient: dockerClient,
	}, nil
}

func (c *nativeClient) ImageForName(name string) (*Image, error) {
	images, err := c.ImagesForRunningContainers()
	if err != nil {
		return nil, err
	}
	for _, image := range images {
		if image.Name == name {
			c.logger.Debugf("found image for name: %s - %v", name, image)
			return &image, nil
		}
	}
	return nil, constants.ErrNoSuchImage
}

func (c *nativeClient) ImagesForRunningContainers() ([]Image, error) {
	containers, err := c.dockerClient.ContainerList(ctx, listOptions)
	if err != nil {
		return nil, err
	}
	var images []Image
	for _, container := range containers {
		images = append(images, fromContainer(container))
	}
	c.logger.Debugf("fetched images for running containers: %v", images)
	return images, nil
}
