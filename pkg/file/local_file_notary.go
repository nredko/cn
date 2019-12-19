package file

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"

	"github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/logger"
	"github.com/codenotary/ctrlt/pkg/notary"
)

type LocalFileNotary struct {
	logger logger.Logger
	notary notary.Notary
}

func NewLocalFileNotary() (FileNotary, error) {
	repository, err := di.Lookup(constants.Notary)
	if err != nil {
		return nil, err
	}
	log, err := di.Lookup(constants.Logger)
	if err != nil {
		return nil, err
	}
	return &LocalFileNotary{
		logger: log.(logger.Logger),
		notary: repository.(notary.Notary),
	}, nil
}

func (n *LocalFileNotary) Notarize(path string, status string) (*notary.Notarization, error) {
	hash, err := n.HashForFile(path)
	if err != nil {
		return nil, err
	}
	notarization, err := n.notary.Notarize(hash, status)
	if err != nil {
		return nil, err
	}
	n.logger.Debugf("notarized %s: %v", path, notarization)
	return notarization, nil
}

func (n *LocalFileNotary) Authenticate(path string) (*notary.Notarization, error) {
	hash, err := n.HashForFile(path)
	if err != nil {
		return nil, err
	}
	notarization, err := n.notary.Authenticate(hash)
	if err != nil {
		return nil, err
	}
	n.logger.Debugf("authenticated %s: %v", path, notarization)
	return notarization, nil
}

func (n *LocalFileNotary) History(path string) ([]*notary.Notarization, error) {
	hash, err := n.HashForFile(path)
	if err != nil {
		return nil, err
	}
	history, err := n.notary.History(hash)
	if err != nil {
		return nil, err
	}
	n.logger.Debugf("notarization history %s: %v", path, history)
	return history, nil
}

func (n *LocalFileNotary) HashForFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	h := sha256.New()
	if _, err = io.Copy(h, file); err != nil {
		return "", err
	}
	checksum := h.Sum(nil)
	return hex.EncodeToString(checksum), nil
}
