package repository

import (
	"context"
	"database/sql"

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
}

func NewRepository(postgresDB *sql.DB, redisClient *redis.Client) *Repository {
	return &Repository{
		UrlShortenerPostgres: NewUrlShortenerPostgreSQL(postgresDB),
		UrlShortenerRedis:    NewUrlShortenerRedis(redisClient),
	}
}
