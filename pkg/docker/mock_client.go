package docker

type mockClient struct{}

func NewMockClient() (Client, error) {
	return &mockClient{}, nil
}

func (m mockClient) ImageForName(name string) (*Image, error) {
	return &Image{
		Name: "name",
		Hash: "hash",
	}, nil
}

func (m mockClient) ImagesForRunningContainers() ([]Image, error) {
	return []Image{
		{
			Name: "name",
			Hash: "hash",
		},
	}, nil
}
