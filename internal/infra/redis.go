package infra

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"exusiai.dev/livehouse/internal/config"
)

func Redis(conf *config.Config) (*redis.Client, error) {
	u, err := redis.ParseURL(conf.RedisURL)
	if err != nil {
		return nil, err
	}

	// Open a Redis Client
	client := redis.NewClient(u)

	// check redis connection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	ping := client.Ping(ctx)
	if ping.Err() != nil {
		return nil, ping.Err()
	}

	return client, nil
}
