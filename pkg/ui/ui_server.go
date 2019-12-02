package ui

import (
	"net/http"
)

type uiServer struct{}

func NewUiServer() (Server, error) {
	return &uiServer{}, nil
}

func (s *uiServer) Start() error {
	http.Handle("/", http.FileServer(http.Dir("ui")))
	return nil
}
