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
	"github.com/gorilla/mux"
)

const (
	timeout = 2 * time.Second

	shrotUrl = "/{uri}"
)

// @Summary Get short link
// @Tags Operations with url
// @Description  get short link
// @ModuleID GetShortLink
// @Accept  json
// @Produce  json
// @Param request body model.LinkRequest true "url"
// @Success 200 {object} string
// @Failure 400 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router / [post]
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

// @Summary Get long link
// @Tags Operations with url
// @Description  get long link
// @ModuleID GetLongLink
// @Accept  json
// @Produce  json
// @Param uri path string true "uri"
// @Success 200 {object} string
// @Failure 400,404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /{uri} [get]
func (h *handler) GetLongLink(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	params := mux.Vars(r)
	shortURI := params["uri"]

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
