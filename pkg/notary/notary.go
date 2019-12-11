package notary

type Notary interface {
	Start() error
	Notarize(hash string, status string) (*Notarization, error)
	Authenticate(hash string) (*Notarization, error)
	AuthenticateBatch(hashes []string) ([]Notarization, error)
	History(hash string) ([]*Notarization, error)
}
