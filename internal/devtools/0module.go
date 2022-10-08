package devtools

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("devtools", fx.Invoke(
		RegisterTracing,
	))
}
