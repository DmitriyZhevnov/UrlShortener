package service

import (
	"context"

	"github.com/DmitriyZhevnov/UrlShortener/internal/repository"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/utils"
)

type urlShortenerService struct {
	postgresRepository repository.UrlShortenerPostgres
	redisRepository    repository.UrlShortenerRedis
	hasher             utils.LinkHasher
}

func NewUrlShortenerSevice(p repository.UrlShortenerPostgres, r repository.UrlShortenerRedis, hasher utils.LinkHasher) *urlShortenerService {
	return &urlShortenerService{
		postgresRepository: p,
		redisRepository:    r,
		hasher:             hasher,
	}
}

func (s *urlShortenerService) Get(ctx context.Context, longLink string) (string, error) {
	_, err := s.hasher.IsValidLink(longLink)
	if err != nil {
		// TODO
		return "", err
	}

	return s.postgresRepository.Get(ctx, longLink)
}
