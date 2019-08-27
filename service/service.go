package service

import (
	"net/http"

	"github.com/ecadlabs/tezos-streamer/errors"
	"github.com/ecadlabs/tezos-streamer/middleware"
	"github.com/ecadlabs/tezos-streamer/streamer"
	"github.com/ecadlabs/tezos-streamer/utils"
	"github.com/gorilla/mux"
)

type Service struct {
	handler *Handler
}

func NewService(str *streamer.Streamer) (*Service, error) {
	return &Service{
		handler: &Handler{
			streamer: str,
		},
	}, nil
}

func (s *Service) NewAPIHandler() http.Handler {
	m := mux.NewRouter()
	// m.Use((&middleware.Logging{}).Handler)
	m.Use((&middleware.Recover{}).Handler)

	m.Methods("GET").Path("/subscribe").HandlerFunc(s.handler.Subscribe)

	m.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.JSONError(w, errors.ErrResourceNotFound)
	})

	return m
}
