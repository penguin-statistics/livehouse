package httpsvr

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
	"github.com/rs/xid"

	"github.com/penguin-statistics/livehouse/internal/config"
)

type Server struct {
	*fiber.App
}

func Create(conf *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:                 "penguin-livehouse",
		ServerHeader:            "penguin-livehouse",
		ProxyHeader:             fiber.HeaderXForwardedFor,
		EnableTrustedProxyCheck: true,
		TrustedProxies: []string{
			"::1",
			"127.0.0.1",
		},
		Immutable: true,
	})

	app.Use(favicon.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, DELETE, OPTIONS",
		AllowHeaders:     "Content-Type, Authorization, X-Requested-With, X-Penguin-Variant, sentry-trace",
		ExposeHeaders:    "Content-Type, X-Penguin-Request-ID",
		AllowCredentials: true,
	}))

	app.Use(helmet.New(helmet.Config{
		HSTSMaxAge:         31356000,
		HSTSPreloadEnabled: true,
		ReferrerPolicy:     "strict-origin-when-cross-origin",
		PermissionPolicy:   "interest-cohort=()",
	}))

	if conf.DevMode {
		app.Use(pprof.New())
	}

	if conf.LimiterEnabled {
		app.Use(limiter.New(limiter.Config{
			Expiration: time.Minute,
			Max:        10,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
					"code":    "TOO_MANY_REQUESTS",
					"message": "Your client is sending requests too frequently. Consult X-RateLimit-Limit and X-RateLimit-Remaining headers for details on your current rate limitation status.",
				})
			},
		}))
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
