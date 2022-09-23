package repository

import (
	"context"
	"fmt"

	"github.com/DmitriyZhevnov/UrlShortener/internal/apperror"
	"github.com/go-redis/redis/v8"
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
	return s.client.Get(ctx, longLink).Result()
}

func (s *urlShortenerRedis) Post(ctx context.Context, longLink, shortLink string) error {
	_, err := s.client.Set(ctx, longLink, shortLink, 0).Result()
	if err != nil {
		return apperror.NewInternalServerError(fmt.Sprintf("failed to save into redis due to error: %v", err))
	}

	return nil
}
