package service

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

var (
	ErrRedisNotReachable = errors.New("redis not reachable")
	ErrNATSNotReachable  = errors.New("nats not reachable")
)

type Health struct {
	Redis *redis.Client
	NATS  *nats.Conn
}

func NewHealth(redis *redis.Client, nats *nats.Conn) *Health {
	return &Health{
		Redis: redis,
		NATS:  nats,
	}
}

func (s *Health) Ping(ctx context.Context) error {
	if err := s.Redis.Ping(ctx).Err(); err != nil {
		return errors.Wrap(ErrRedisNotReachable, err.Error())
	}

	// nats does automatic ping for 20 seconds interval (configurated at infra/nats.go)
	status := s.NATS.Status()
	if status != nats.CONNECTED && status != nats.DRAINING_PUBS && status != nats.DRAINING_SUBS {
		return errors.Wrap(ErrNATSNotReachable, status.String())
	}

	return nil
}
