package notary

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/codenotary/immudb/pkg/client"

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
	n := storedNotarization{}
	if err = json.Unmarshal(response.Value, &n); err != nil {
		return nil, err
	}
	r.logger.Debugf("get %s - %s @ %v", hash, response.Index, n)
	return &Notarization{
		Hash:      n.Hash,
		Status:    n.Status,
		Meta:      n.Meta,
		StoreMeta: NewStoreMeta(response.Index),
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
		n := storedNotarization{}
		if err = json.Unmarshal(item.Value, &n); err != nil {
			return nil, err
		}
		notarizations = append(notarizations, &Notarization{
			Hash:      n.Hash,
			Status:    n.Status,
			Meta:      n.Meta,
			StoreMeta: NewStoreMeta(item.Index),
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
	for _, response := range batchResponse.Items {
		if len(response.Value) == 0 {
			notarizations = append(notarizations, *UnknownNotarization)
		} else {
			n := storedNotarization{}
			if err = json.Unmarshal(response.Value, &n); err != nil {
				return nil, err
			}
			notarizations = append(notarizations, Notarization{
				Hash:      n.Hash,
				Status:    n.Status,
				Meta:      n.Meta,
				StoreMeta: NewStoreMeta(response.Index),
			})
		}
	}
	r.logger.Debugf("get-batch %v - %v", hashes, notarizations)
	return notarizations, nil
}

func (r *immuNotary) Notarize(hash string, status string, meta Meta) (*Notarization, error) {
	key := bytes.NewReader([]byte(hash))
	value, err := json.Marshal(&storedNotarization{
		Hash:   hash,
		Status: status,
	})
	if err != nil {
		return nil, err
	}
	response, err := r.immuClient.Set(key, bytes.NewReader(value))
	if err != nil {
		return nil, err
	}
	r.logger.Debugf("create %s - %s @ %d", hash, status, response.Index)
	return &Notarization{
		Hash:      hash,
		Status:    status,
		Meta:      meta,
		StoreMeta: NewStoreMeta(response.Index),
	}, nil
}

func NewStoreMeta(index uint64) StoreMeta {
	return StoreMeta{"index": index}
}
