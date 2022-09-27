package handler

import (
	"net/http"

	"github.com/DmitriyZhevnov/UrlShortener/internal/apperror"
	"github.com/DmitriyZhevnov/UrlShortener/internal/service"
	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	Register(router *httprouter.Router)
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

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/", apperror.MiddleWare(h.GetShortLink))
	router.HandlerFunc(http.MethodGet, shrotUrl, apperror.MiddleWare(h.GetLongLink))
}
