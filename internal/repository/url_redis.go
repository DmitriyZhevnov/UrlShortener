package repository

import (
	"context"

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
	return "", nil
}
