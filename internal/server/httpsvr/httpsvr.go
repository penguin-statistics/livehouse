package httpsvr

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/xid"

	"github.com/penguin-statistics/livehouse/internal/config"
)

type Server struct {
	*fiber.App
}

func Create(conf *config.Config) *fiber.App {
	app := fiber.New()

	if conf.DevMode {
		app.Use(pprof.New())
	}

	app.Use(requestid.New(requestid.Config{
		Header:     "X-Penguin-Request-ID",
		ContextKey: "requestid",
		Generator: func() string {
			return xid.New().String()
		},
	}))

	return app
}
