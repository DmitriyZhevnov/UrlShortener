package handler

import (
	"github.com/DmitriyZhevnov/UrlShortener/internal/apperror"
	"github.com/DmitriyZhevnov/UrlShortener/internal/service"
	"github.com/gorilla/mux"
)

type Handler interface {
	Register(router *mux.Router)
}

type handler struct {
	service *service.Service
	domain  string
}

func NewHandler(service *service.Service, domain string) Handler {
	return &handler{
		service: service,
		domain:  domain,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/", apperror.MiddleWare(h.GetShortLink)).Methods("POST")
	router.HandleFunc(shrotUrl, apperror.MiddleWare(h.GetLongLink)).Methods("GET")
}
