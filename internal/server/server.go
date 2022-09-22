package server

import (
	"context"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(router *httprouter.Router) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         ":8080",
			Handler:      router,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
