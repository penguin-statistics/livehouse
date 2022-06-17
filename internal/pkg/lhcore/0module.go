package lhcore

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("lhcore", fx.Provide(
		NewDropSet,
	))
}
