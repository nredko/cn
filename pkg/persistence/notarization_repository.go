package persistence

type NotarizationRepository interface {
	Start() error
	GetNotarizationForHash(hash string) (*Notarization, error)
	GetNotarizationHistoryForHash(hash string) ([]*Notarization, error)
	GetNotarizationsForHashes(hashes []string) ([]Notarization, error)
	CreateNotarization(hash string, status string) (*Notarization, error)
}
