package service

import (
	"context"

	"github.com/DmitriyZhevnov/UrlShortener/internal/repository"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/logging"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/utils"
)

type Service struct {
	UrlShortener
}

type UrlShortener interface {
	GetShortLink(ctx context.Context, longLink string) (string, error)
	GetLongLink(ctx context.Context, shortLink string) (string, error)
}

func NewService(logger logging.Logger, repos *repository.Repository, hasher utils.LinkHasher) *Service {
	return &Service{
		UrlShortener: NewUrlShortenerSevice(logger, repos.UrlShortenerPostgres, repos.UrlShortenerRedis, hasher),
	}
}
