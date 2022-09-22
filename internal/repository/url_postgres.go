package repository

import (
	"context"
	"database/sql"
)

type urlShortenerPostgres struct {
	db *sql.DB
}

func NewUrlShortenerPostgreSQL(db *sql.DB) *urlShortenerPostgres {
	return &urlShortenerPostgres{
		db: db,
	}
}

func (s *urlShortenerPostgres) Get(ctx context.Context, longLink string) (string, error) {
	return "", nil
}
