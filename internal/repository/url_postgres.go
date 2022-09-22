package repository

import (
	"context"
	"database/sql"
	"errors"
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
	q := `
		SELECT short_link FROM public.link WHERE long_link = $1
	`

	var shortLink string
	err := s.db.QueryRow(q, longLink).Scan(&shortLink)
	if err != nil {
		// TODO
		return "", errors.New("not found")
	}

	return shortLink, nil
}

func (s *urlShortenerPostgres) Post(ctx context.Context, longLink, shortLink string) error {
	q := `
        INSERT INTO link
            (long_link, short_link)
        VALUES
               ($1, $2)
    `
	if err := s.db.QueryRow(q, longLink, shortLink); err.Err() != nil {
		// TODO
		return err.Err()
	}

	return nil
}
