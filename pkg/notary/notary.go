package notary

//go:generate mockgen -source=../notary/notary.go -destination=../mocks/mock_notary.go -package=mocks

type Notary interface {
	Start() error
	Notarize(hash string, status string, meta Meta) (*Notarization, error)
	Authenticate(hash string) (*Notarization, error)
	AuthenticateBatch(hashes []string) ([]Notarization, error)
	History(hash string) ([]*Notarization, error)
}
