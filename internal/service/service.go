package service

import (
	"context"

	"github.com/DmitriyZhevnov/UrlShortener/internal/repository"
)

type Service struct {
	UrlShortener
}

type UrlShortener interface {
	Get(ctx context.Context, longLink string) (string, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UrlShortener: NewUrlShortenerSevice(repos.UrlShortenerPostgres, repos.UrlShortenerRedis),
	}
}
