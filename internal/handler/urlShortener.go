package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/DmitriyZhevnov/UrlShortener/internal/apperror"
	"github.com/DmitriyZhevnov/UrlShortener/internal/model"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/response"
	"github.com/julienschmidt/httprouter"
)

const (
	timeout = 2 * time.Second

	shrotUrl = "/:uri"
)

func (h *handler) GetShortLink(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	request := model.LinkRequest{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		return apperror.NewBadRequestError("invalid json")
	}

	shortLink, err := h.service.GetShortLink(ctx, request.Url)
	if err != nil {
		return err
	}

	response.SendResponse(w, 200, shortLink)
	return nil
}

func (h *handler) GetLongLink(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	shortURI := httprouter.ParamsFromContext(r.Context()).ByName("uri")

	var b strings.Builder
	b.WriteString(h.domain)
	b.WriteString("/")
	b.WriteString(shortURI)

	longLink, err := h.service.GetLongLink(ctx, b.String())
	if err != nil {
		return err
	}

	response.SendResponse(w, 200, longLink)
	return nil
}
