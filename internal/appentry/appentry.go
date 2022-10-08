package appentry

import (
	"time"

	"go.uber.org/fx"

	"github.com/penguin-statistics/livehouse/internal/config"
	"github.com/penguin-statistics/livehouse/internal/controller"
	"github.com/penguin-statistics/livehouse/internal/devtools"
	"github.com/penguin-statistics/livehouse/internal/infra"
	"github.com/penguin-statistics/livehouse/internal/pkg/lhcore"
	"github.com/penguin-statistics/livehouse/internal/pkg/logger"
	"github.com/penguin-statistics/livehouse/internal/pkg/wshub"
	"github.com/penguin-statistics/livehouse/internal/repo"
	"github.com/penguin-statistics/livehouse/internal/server/grpcsvr"
	"github.com/penguin-statistics/livehouse/internal/server/httpsvr"
	"github.com/penguin-statistics/livehouse/internal/service"
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
