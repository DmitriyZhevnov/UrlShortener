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
}

func NewHandler(service *service.Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, generateUrl, apperror.MiddleWare(h.GetShortLink))
}
