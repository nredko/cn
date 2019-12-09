package notarization

import (
	"regexp"

	"github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/docker"
	"github.com/codenotary/ctrlt/pkg/logger"
	"github.com/codenotary/ctrlt/pkg/persistence"
)

type dockerNotary struct {
	logger       logger.Logger
	dockerClient docker.Client
	repository   persistence.NotarizationRepository
}

func NewDockerNotary() (ContainerNotary, error) {
	dockerClient, err := di.Lookup(constants.DockerClient)
	if err != nil {
		return nil, err
	}
	repository, err := di.Lookup(constants.NotarizationRepository)
	if err != nil {
		return nil, err
	}
	log, err := di.Lookup(constants.Logger)
	if err != nil {
		return nil, err
	}
	return &dockerNotary{
		logger:       log.(logger.Logger),
		dockerClient: dockerClient.(docker.Client),
		repository:   repository.(persistence.NotarizationRepository),
	}, nil
}

func (n *dockerNotary) ListNotarizedImages(query string) ([]NotarizedImage, error) {
	runningImages, err := n.dockerClient.ImagesForRunningContainers()
	if err != nil {
		return nil, err
	}
	var images []docker.Image
	for _, runningImage := range runningImages {
		match, err := regexp.MatchString(query, runningImage.Name)
		if err != nil {
			return nil, err
		}
		if query == "" || match {
			images = append(images, runningImage)
		}
	}
	var hashes []string
	for _, image := range images {
		hashes = append(hashes, image.Hash)
	}
	notarizations, err := n.repository.GetNotarizationsForHashes(hashes)
	if err != nil {
		return nil, err
	}
	var notarizedImages []NotarizedImage
	for i, image := range images {
		notarizedImages = append(notarizedImages, NotarizedImage{
			Image:        image,
			Notarization: notarizations[i],
		})
	}
	n.logger.Debugf("fetched notarized images: %v", notarizedImages)
	return notarizedImages, nil
}

func (n *dockerNotary) Notarize(hash string, status string) (*persistence.Notarization, error) {
	notarization, err := n.repository.CreateNotarization(hash, status)
	if err != nil {
		return nil, err
	}
	n.logger.Debugf("notarized %s - %s", hash, status)
	return notarization, nil
}

func (n *dockerNotary) NotarizeImageWithName(name string, status string) (*persistence.Notarization, error) {
	image, err := n.dockerClient.ImageForName(name)
	if err != nil {
		return nil, err
	}
	return n.Notarize(image.Hash, status)
}

func (n *dockerNotary) GetNotarizationForHash(hash string) (*persistence.Notarization, error) {
	return n.repository.GetNotarizationForHash(hash)
}

func (n *dockerNotary) GetNotarizationHistoryForHash(hash string) ([]*persistence.Notarization, error) {
	return n.repository.GetNotarizationHistoryForHash(hash)
}

func (n *dockerNotary) GetFirstNotarizationMatchingName(name string) (*persistence.Notarization, error) {
	image, err := n.dockerClient.ImageForName(name)
	if err != nil {
		return nil, err
	}
	return n.GetNotarizationForHash(image.Hash)
}
