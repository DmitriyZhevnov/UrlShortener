package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func NewClient(ctx context.Context) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	_, err = client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return
}
