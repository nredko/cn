package di

type Entry struct {
	Name  string
	Maker func() (interface{}, error)
}
