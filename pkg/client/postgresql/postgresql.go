package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/DmitriyZhevnov/UrlShortener/internal/config"
	repeatable "github.com/DmitriyZhevnov/UrlShortener/pkg/utils"
	_ "github.com/lib/pq"
)

func NewClient(ctx context.Context, maxAttempts int, sc config.Postgresql) (db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", sc.Host, sc.Port,
		sc.Username, sc.Password, sc.Database)
	err = repeatable.DoWithTries(func() error {
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Fatal("error do with tries postgresql")
	}

	return db, nil
}
