package appentry

import (
	"time"

	"go.uber.org/fx"

	"exusiai.dev/livehouse/internal/config"
	"exusiai.dev/livehouse/internal/controller"
	"exusiai.dev/livehouse/internal/devtools"
	"exusiai.dev/livehouse/internal/infra"
	"exusiai.dev/livehouse/internal/pkg/lhcore"
	"exusiai.dev/livehouse/internal/pkg/logger"
	"exusiai.dev/livehouse/internal/pkg/wshub"
	"exusiai.dev/livehouse/internal/repo"
	"exusiai.dev/livehouse/internal/server/grpcsvr"
	"exusiai.dev/livehouse/internal/server/httpsvr"
	"exusiai.dev/livehouse/internal/service"
)

func ProvideOptions() []fx.Option {
	opts := []fx.Option{
		// Misc
		fx.Provide(config.Parse),
		fx.Provide(grpcsvr.Create),
		fx.Provide(httpsvr.Create),
		fx.Provide(wshub.NewHub),

		// Infrastructures
		infra.Module(),

		// Repositories
		repo.Module(),

		// Services
		service.Module(),

		// lhcore
		lhcore.Module(),

		// Global Singleton Inits: Keep those before controllers to ensure they are initialized
		// before controllers are registered as controllers are also fx#Invoke functions which
		// are called in the order of their registration.
		fx.Invoke(logger.Configure),
		fx.Invoke(infra.SentryInit),

		// Controllers
		controller.Module(),

		devtools.Module(),

		// fx Extra Options
		fx.StartTimeout(1 * time.Second),
		fx.StopTimeout(5 * time.Minute),
	}

	return opts
}
