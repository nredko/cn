package api

import (
	"encoding/json"
	"net/http"

	"github.com/codenotary/logger/pkg/logger"

	"github.com/codenotary/ctrlt/pkg/constants"
	"github.com/codenotary/ctrlt/pkg/di"
)

type apiServer struct {
	logger logger.Logger
}

func NewApiServer() (Server, error) {
	log, err := di.Lookup(constants.Logger)
	if err != nil {
		return nil, err
	}
	server := &apiServer{
		logger: log.(logger.Logger),
	}
	return server, nil
}

func (s *apiServer) Start() error {
	http.HandleFunc("/health", s.healthCheck)
	return nil
}

func (s *apiServer) healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.logger.Errorf("unable to encode json: %v", err)
		return
	}
}
