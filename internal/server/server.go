package server

import (
	"context"
	"net/http"
	"time"

	"github.com/DmitriyZhevnov/UrlShortener/internal/config"
	"github.com/julienschmidt/httprouter"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(router *httprouter.Router, cfg config.HTTP) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      router,
			WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
			ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
