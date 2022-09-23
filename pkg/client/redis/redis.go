package redis

import (
	"context"

	"github.com/DmitriyZhevnov/UrlShortener/internal/config"
	"github.com/go-redis/redis/v8"
)

func NewClient(ctx context.Context, sc config.Redis) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     sc.Addr,
		Password: sc.Password,
		DB:       sc.DB,
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return
}
