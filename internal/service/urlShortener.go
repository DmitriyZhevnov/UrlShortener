package service

import (
	"context"
	"fmt"

	"github.com/DmitriyZhevnov/UrlShortener/internal/apperror"
	"github.com/DmitriyZhevnov/UrlShortener/internal/repository"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/logging"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/utils"
	"golang.org/x/sync/errgroup"
)

type urlShortenerService struct {
	log                logging.Logger
	postgresRepository repository.UrlShortenerPostgres
	redisRepository    repository.UrlShortenerRedis
	hasher             utils.LinkHasher
}

func NewUrlShortenerSevice(logger logging.Logger, p repository.UrlShortenerPostgres,
	r repository.UrlShortenerRedis, hasher utils.LinkHasher) *urlShortenerService {
	return &urlShortenerService{
		log:                logger,
		postgresRepository: p,
		redisRepository:    r,
		hasher:             hasher,
	}
}

func (s *urlShortenerService) Get(ctx context.Context, longLink string) (string, error) {
	url, err := s.hasher.IsValidLink(longLink)
	if err != nil {
		s.log.Warn(fmt.Sprintf("invalid link: '%s'", longLink), nil)
		return "", apperror.NewBadRequestError("invalid link")
	}

	shortLink, err := s.redisRepository.Get(ctx, longLink)
	if err == nil {
		s.log.Info(fmt.Sprintf("short link for '%s' has been found in redis.", longLink), nil)
		return shortLink, nil
	}

	shortLink, err = s.postgresRepository.Get(ctx, longLink)
	if err == nil {
		s.log.Info(fmt.Sprintf("short link for '%s' has been found in postgres.", longLink), nil)
		return shortLink, nil
	}

	shortLink = s.hasher.GenerateShortLink(url)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return s.postgresRepository.Post(ctx, longLink, shortLink)
	})

	g.Go(func() error {
		return s.redisRepository.Post(ctx, longLink, shortLink)
	})

	if err = g.Wait(); err != nil {
		s.log.Error(fmt.Sprintf("error occurred while saving to db: %s", err.Error()), nil)
		return "", err
	}

	s.log.Info(fmt.Sprintf("short link for '%s' has been successfully saved in the database.", longLink), nil)
	return shortLink, nil
}
