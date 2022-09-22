package service

import (
	"context"

	"github.com/DmitriyZhevnov/UrlShortener/internal/repository"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/utils"
	"golang.org/x/sync/errgroup"
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
	url, err := s.hasher.IsValidLink(longLink)
	if err != nil {
		// TODO
		return "", err
	}

	shortLink, err := s.redisRepository.Get(ctx, longLink)
	if err == nil {
		return shortLink, nil
	}

	shortLink, err = s.postgresRepository.Get(ctx, longLink)
	if err == nil {
		return shortLink, nil
	}

	shortLink = s.hasher.HashURI(url)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return s.postgresRepository.Post(ctx, longLink, shortLink)
	})

	g.Go(func() error {
		return s.redisRepository.Post(ctx, longLink, shortLink)
	})

	if err = g.Wait(); err != nil {
		// TODO
		return "", err
	}

	return shortLink, nil
}
