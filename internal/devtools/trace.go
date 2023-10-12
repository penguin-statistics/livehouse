package devtools

import (
	"context"
	"os"
	"runtime/trace"

	"go.uber.org/fx"

	"exusiai.dev/livehouse/internal/config"
)

func RegisterTracing(conf *config.Config, lc fx.Lifecycle) error {
	if conf.DevMode {
		f, err := os.Create("/tmp/trace.out")
		if err != nil {
			return err
		}

		err = trace.Start(f)
		if err != nil {
			return err
		}

		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				trace.Stop()
				return nil
			},
		})
	}

	return nil
}
