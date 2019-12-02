package di

type Initializing interface {
	Start() error
}

type Terminating interface {
	Stop() error
}
