package service

import (
	"context"
	"net"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/penguin-statistics/livehouse/internal/config"
	"github.com/penguin-statistics/livehouse/internal/pkg/async"
)

func run(grpcserv *grpc.Server, httpserv *fiber.App, conf *config.Config, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", conf.GRPCAddress)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to listen")
			}

			go func() {
				if err := grpcserv.Serve(ln); err != nil {
					log.Error().Err(err).Msg("server terminated unexpectedly")
				}
			}()

			go func() {
				if err := httpserv.Listen(conf.HTTPAddress); err != nil {
					log.Error().Err(err).Msg("server terminated unexpectedly")
				}
			}()

			log.Info().Msgf("started gRPC server on %s", conf.GRPCAddress)
			log.Info().Msgf("started HTTP server on %s", conf.HTTPAddress)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if conf.DevMode {
				return nil
			}

			return async.WaitAll(
				async.Errable(func() error {
					log.Info().Msg("grpc server is shutting down")

					grpcserv.GracefulStop()

					return nil
				}),
				async.Errable(func() error {
					log.Info().Msg("http server is shutting down")

					return httpserv.Shutdown()
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
