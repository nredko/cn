package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/di"
	"github.com/codenotary/ctrlt/pkg/logger"
	"github.com/codenotary/ctrlt/pkg/notarization"
)

type apiServer struct {
	logger logger.Logger
	notary notarization.ContainerNotary
}

func NewApiServer() (Server, error) {
	notary, err := di.Lookup(constants.Notary)
	if err != nil {
		return nil, err
	}
	log, err := di.Lookup(constants.Logger)
	if err != nil {
		return nil, err
	}
	server := &apiServer{
		logger: log.(logger.Logger),
		notary: notary.(notarization.ContainerNotary),
	}
	return server, nil
}

func (s *apiServer) Start() error {
	http.HandleFunc("/containers", s.listContainers)
	http.HandleFunc("/history", s.containerHistory)
	http.HandleFunc("/notarize", s.notarize)
	http.HandleFunc("/bulk-notarize", s.bulkNotarize)
	http.HandleFunc("/health", s.healthCheck)
	return nil
}

func (s *apiServer) listContainers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query().Get("query")
	notarizedImages, err := s.notary.ListNotarizedImages(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.logger.Errorf("notarization failed: %v", err)
		return
	}
	if err := json.NewEncoder(w).Encode(notarizedImages); err != nil {
		s.logger.Errorf("unable to encode json: %v", err)
		return
	}
}

func (s *apiServer) notarize(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	status := r.URL.Query().Get("status")
	w.Header().Set("Content-Type", "application/json")
	if hash == "" || status == "" {
		w.WriteHeader(http.StatusBadRequest)
		s.logger.Warningf("hash or status parameters missing")
		return
	}
	if _, err := s.notary.Notarize(hash, status); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.logger.Errorf("notarization failed: %v", err)
		return
	}
	if err := json.NewEncoder(w).Encode("ok"); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		s.logger.Errorf("unable to encode json: %v", err)
		return
	}
	s.logger.Infof("notarized %s with status %s\n", hash, status)
}

func (s *apiServer) bulkNotarize(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	w.Header().Set("Content-Type", "application/json")
	if status == "" {
		w.WriteHeader(http.StatusBadRequest)
		s.logger.Warningf("status parameter missing")
		return
	}
	query := r.URL.Query().Get("query")
	notarizedImages, err := s.notary.ListNotarizedImages(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.logger.Errorf("notarization failed: %v", err)
		return
	}
	for _, image := range notarizedImages {
		if _, err := s.notary.Notarize(image.Image.Hash, status); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.logger.Errorf("notarization failed: %v", err)
			return
		}
	}
	if err := json.NewEncoder(w).Encode("ok"); err != nil {
		s.logger.Errorf("unable to encode json: %v", err)
		return
	}
	s.logger.Infof("notarized %v with status %s\n", notarizedImages, status)
}

func (s *apiServer) healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode("ok"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.logger.Errorf("unable to encode json: %v", err)
		return
	}
}

func (s *apiServer) containerHistory(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	w.Header().Set("Content-Type", "application/json")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		s.logger.Warningf("hash parameter missing")
		return
	}
	notarizationHistory, err := s.notary.GetNotarizationHistoryForHash(hash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.logger.Errorf("notarization failed: %v", err)
		return
	}
	if err := json.NewEncoder(w).Encode(notarizationHistory); err != nil {
		s.logger.Errorf("unable to encode json: %v", err)
		return
	}
}
