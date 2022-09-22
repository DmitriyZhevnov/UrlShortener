package repository

import (
	"context"
	"fmt"

	"github.com/DmitriyZhevnov/UrlShortener/internal/apperror"
	"gopkg.in/redis.v3"
)

type urlShortenerRedis struct {
	client *redis.Client
}

func NewUrlShortenerRedis(client *redis.Client) *urlShortenerRedis {
	return &urlShortenerRedis{
		client: client,
	}
}

func (s *urlShortenerRedis) Get(ctx context.Context, longLink string) (string, error) {
	return s.client.Get(longLink).Result()
}

func (s *urlShortenerRedis) Post(ctx context.Context, longLink, shortLink string) error {
	_, err := s.client.Set(longLink, shortLink, 0).Result()
	if err != nil {
		return apperror.NewInternalServerError(fmt.Sprintf("failed to save into redis due to error: %v", err))
	}

	return nil
}
