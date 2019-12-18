package notary

import (
	"bytes"
	"io"

	"github.com/codenotary/immustore/pkg/client"

	"github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/logger"
)

type immuNotary struct {
	logger     logger.Logger
	immuClient *client.ImmuClient
}

func NewImmuNotary() (Notary, error) {
	log, err := di.Lookup(constants.Logger)
	if err != nil {
		return nil, err
	}
	immuClient, err := di.Lookup(constants.ImmuClient)
	if err != nil {
		return nil, err
	}
	return &immuNotary{
		logger:     log.(logger.Logger),
		immuClient: immuClient.(*client.ImmuClient),
	}, nil
}

func (r *immuNotary) Start() error {
	return r.immuClient.Connect()
}

func (r *immuNotary) Stop() error {
	return r.immuClient.Disconnect()
}

func (r *immuNotary) Authenticate(hash string) (*Notarization, error) {
	response, err := r.immuClient.Get(bytes.NewReader([]byte(hash)))
	if err != nil {
		return UnknownNotarization, nil
	}
	status := string(response.Value)
	r.logger.Debugf("get %s - %s @ %d", hash, response.Index, status)
	return &Notarization{
		Hash:   hash,
		Status: status,
		Index:  response.Index,
	}, nil
}

func (r *immuNotary) History(hash string) ([]*Notarization, error) {
	response, err := r.immuClient.History(bytes.NewReader([]byte(hash)))
	if err != nil {
		return nil, err
	}
	r.logger.Debugf("history %s - %v", hash, response.Items)
	var notarizations []*Notarization
	for _, item := range response.Items {
		status := string(item.Value)
		notarizations = append(notarizations, &Notarization{
			Hash:   hash,
			Status: status,
			Index:  item.Index,
		})
	}
	return notarizations, nil
}

func (r *immuNotary) AuthenticateBatch(hashes []string) ([]Notarization, error) {
	var readers []io.Reader
	for _, hash := range hashes {
		readers = append(readers, bytes.NewReader([]byte(hash)))
	}
	batchResponse, err := r.immuClient.GetBatch(readers)
	if err != nil {
		return nil, err
	}
	var notarizations []Notarization
	for i, response := range batchResponse.GetResponses {
		if len(response.Value) == 0 {
			notarizations = append(notarizations, *UnknownNotarization)
		} else {
			notarizations = append(notarizations, Notarization{
				Hash:   hashes[i],
				Status: string(response.Value),
				Index:  response.Index,
			})
		}
	}
	r.logger.Debugf("get-batch %v - %v", hashes, notarizations)
	return notarizations, nil
}

func (r *immuNotary) Notarize(hash string, status string) (*Notarization, error) {
	key := bytes.NewReader([]byte(hash))
	value := bytes.NewReader([]byte(status))
	response, err := r.immuClient.Set(key, value)
	if err != nil {
		return nil, err
	}
	r.logger.Debugf("create %s - %s @ %d", hash, status, response.Index)
	return &Notarization{
		Hash:   hash,
		Status: status,
		Index:  response.Index,
	}, nil
}
