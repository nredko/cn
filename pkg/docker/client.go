package docker

type Client interface {
	ImageForName(name string) (*Image, error)
	ImagesForRunningContainers() ([]Image, error)
}
