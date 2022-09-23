package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/DmitriyZhevnov/UrlShortener/internal/apperror"
	"github.com/DmitriyZhevnov/UrlShortener/pkg/client/postgresql"
	"github.com/jackc/pgconn"
)

type urlShortenerPostgres struct {
	client postgresql.Client
}

func NewUrlShortenerPostgreSQL(client postgresql.Client) *urlShortenerPostgres {
	return &urlShortenerPostgres{
		client: client,
	}
}

func (s *urlShortenerPostgres) Get(ctx context.Context, longLink string) (string, error) {
	q := `
		SELECT short_link FROM public.link WHERE long_link = $1
	`

	var shortLink string
	err := s.client.QueryRow(ctx, q, longLink).Scan(&shortLink)
	if err != nil {
		return "", err
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
	_, err := s.client.Query(ctx, q, longLink, shortLink)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := apperror.NewInternalServerError(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return newErr
		}
		return apperror.NewInternalServerError(fmt.Sprintf("failed to save into postgres due to error: %v", err))
	}

	return nil
}
