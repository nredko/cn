package notary

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/codenotary/objects/pkg/object"
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

func (r *immuNotary) Authenticate(object *object.Object) (*Notarization, error) {
	response, err := r.immuClient.Get(bytes.NewReader([]byte(object.Digest.Encoded())))
	if err != nil {
		return UnknownNotarization, nil
	}
	n := storedNotarization{}
	if err = json.Unmarshal(response.Value, &n); err != nil {
		return nil, err
	}
	r.logger.Debugf("get %s - %s @ %v", object.Digest.Encoded(), response.Index, n)
	return &Notarization{
		Status:    n.Status,
		Object:    n.Object,
		StoreMeta: NewStoreMeta(response.Index),
	}, nil
}

func (r *immuNotary) History(object *object.Object) ([]*Notarization, error) {
	response, err := r.immuClient.History(bytes.NewReader([]byte(object.Digest.Encoded())))
	if err != nil {
		return nil, err
	}
	r.logger.Debugf("history %s - %v", object.Digest.Encoded(), response.Items)
	var notarizations []*Notarization
	for _, item := range response.Items {
		n := storedNotarization{}
		if err = json.Unmarshal(item.Value, &n); err != nil {
			return nil, err
		}
		notarizations = append(notarizations, &Notarization{
			Status:    n.Status,
			Object:    n.Object,
			StoreMeta: NewStoreMeta(item.Index),
		})
	}
	return notarizations, nil
}

func (r *immuNotary) AuthenticateBatch(objects []*object.Object) ([]Notarization, error) {
	var readers []io.Reader
	for _, o := range objects {
		readers = append(readers, bytes.NewReader([]byte(o.Digest.Encoded())))
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
				Status:    n.Status,
				Object:    n.Object,
				StoreMeta: NewStoreMeta(response.Index),
			})
		}
	}
	r.logger.Debugf("get-batch %v - %v", objects, notarizations)
	return notarizations, nil
}

func (r *immuNotary) Notarize(object *object.Object, status string) (*Notarization, error) {
	key := bytes.NewReader([]byte(object.Digest.Encoded()))
	value, err := json.Marshal(&storedNotarization{
		Object: object,
		Status: status,
	})
	if err != nil {
		return nil, err
	}
	response, err := r.immuClient.Set(key, bytes.NewReader(value))
	if err != nil {
		return nil, err
	}
	r.logger.Debugf("create %s - %s @ %d", object.Digest.Encoded(), status, response.Index)
	return &Notarization{
		Status:    status,
		Object:    object,
		StoreMeta: NewStoreMeta(response.Index),
	}, nil
}

func NewStoreMeta(index uint64) StoreMeta {
	return StoreMeta{"index": index}
}
