package service

import (
	"context"

	"github.com/DmitriyZhevnov/UrlShortener/internal/repository"
)

type urlShortenerService struct {
	postgresRepository repository.UrlShortenerPostgres
	redisRepository    repository.UrlShortenerRedis
}

func NewUrlShortenerSevice(p repository.UrlShortenerPostgres, r repository.UrlShortenerRedis) *urlShortenerService {
	return &urlShortenerService{
		postgresRepository: p,
		redisRepository:    r,
	}
}

func (s *urlShortenerService) Get(ctx context.Context, longLink string) (string, error) {
	return s.postgresRepository.Get(ctx, longLink)
}
