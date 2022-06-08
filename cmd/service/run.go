package service

import (
	"context"
	"net"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/penguin-statistics/livehouse/internal/config"
	"github.com/penguin-statistics/livehouse/internal/pkg/async"
)

func run(serv *grpc.Server, conf *config.Config, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen(conf.Network, conf.Address)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to listen")
			}

			go func() {
				if err := serv.Serve(ln); err != nil {
					log.Error().Err(err).Msg("server terminated unexpectedly")
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if conf.DevMode {
				return nil
			}

			return async.WaitAll(
				async.Errable(func() error {
					log.Info().Msg("grpc server is shutting down")

					ch := make(chan struct{})
					go func() {
						serv.GracefulStop()
						close(ch)
					}()

					select {
					case <-ch:
						log.Info().Msg("grpc server shutdown complete")
					case <-time.After(time.Second * 10):
						log.Error().Msg("grpc server shutdown timed out. forcefully shutting down...")
						serv.Stop()
					}

					return nil
				}),
				async.Errable(func() error {
					flushed := sentry.Flush(time.Second * 10)
					if !flushed {
						return errors.New("sentry flush timeout, some events may be lost")
					}
					return nil
				}),
			)
		},
	})
}
