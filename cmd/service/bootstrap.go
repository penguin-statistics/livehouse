package service

import (
	"go.uber.org/fx"

	"github.com/penguin-statistics/livehouse/internal/appentry"
)

func Bootstrap() {
	opts := []fx.Option{}
	opts = append(opts, appentry.ProvideOptions()...)
	opts = append(opts, fx.Invoke(run))

	app := fx.New(opts...)

	app.Run() // blocks
}
