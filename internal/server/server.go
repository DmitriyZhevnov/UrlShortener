package server

import (
	"context"
	"net/http"
	"time"

	"github.com/DmitriyZhevnov/UrlShortener/internal/config"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/logging"
	"github.com/gorilla/mux"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(router *mux.Router, cfg config.HTTP) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      router,
			WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
			ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		},
	}
}

func (s *Server) Run(logger logging.Logger) error {
	logger.Info("starting server.", nil)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
