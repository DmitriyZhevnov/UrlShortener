package repository

import (
	"context"

	"github.com/DmitriyZhevnov/UrlShortener/pkg/client/postgresql"
	"gopkg.in/redis.v3"
)

type Repository struct {
	UrlShortenerPostgres
	UrlShortenerRedis
}

type UrlShortenerPostgres interface {
	Get(ctx context.Context, longLink string) (string, error)
	Post(ctx context.Context, longLink, shortLink string) error
}

type UrlShortenerRedis interface {
	Get(ctx context.Context, longLink string) (string, error)
	Post(ctx context.Context, longLink, shortLink string) error
}

func NewRepository(client postgresql.Client, redisClient *redis.Client) *Repository {
	return &Repository{
		UrlShortenerPostgres: NewUrlShortenerPostgreSQL(client),
		UrlShortenerRedis:    NewUrlShortenerRedis(redisClient),
	}
}
