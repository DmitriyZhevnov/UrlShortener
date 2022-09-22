package redis

import (
	"context"

	"gopkg.in/redis.v3"
)

func NewClient(ctx context.Context) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	_, err = client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return
}
