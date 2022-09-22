package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DmitriyZhevnov/UrlShortener/internal/apperror"
	"github.com/DmitriyZhevnov/UrlShortener/internal/model"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/response"
)

const (
	generateUrl = "/generate"
)

func (h *handler) GetShortLink(w http.ResponseWriter, r *http.Request) error {
	request := model.LinkRequest{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		return apperror.NewBadRequestError("invalid json")
	}

	shortLink, err := h.service.Get(r.Context(), request.Url)
	if err != nil {
		return err
	}

	response.SendResponse(w, 200, shortLink)
	return nil
}
