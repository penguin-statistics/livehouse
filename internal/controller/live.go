package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"exusiai.dev/livehouse/internal/constant"
	"exusiai.dev/livehouse/internal/service"
)

type LiveDeps struct {
	fx.In

	Service *service.Live
}

type Live struct {
	LiveDeps
}

func RegisterLive(app *fiber.App, deps LiveDeps) {
	ctrler := &Live{
		LiveDeps: deps,
	}

	app.Use("/live", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/live", ctrler.Handle)
}

func (c *Live) Handle(ctx *fiber.Ctx) error {
	f := websocket.New(func(conn *websocket.Conn) {
		if conn.Subprotocol() != constant.LiveWebSocketSubprotocol {
			log.Warn().Str("subprotocol", conn.Subprotocol()).Msg("invalid subprotocol")
			err := conn.WriteMessage(websocket.TextMessage, []byte("invalid subprotocol: expect subprotocol "+constant.LiveWebSocketSubprotocol))
			if err != nil {
				log.Error().Err(err).Msg("failed to write message")
			}

			err = conn.WriteControl(websocket.CloseMessage, []byte("invalid subprotocol: expect subprotocol "+constant.LiveWebSocketSubprotocol), time.Now().Add(time.Second*10))
			if err != nil {
				log.Error().Err(err).Msg("failed to write invalid subprotocol message")
			}
			conn.Close()
			return
		}

		c.Service.Handle(conn)
	}, websocket.Config{
		Subprotocols:      []string{constant.LiveWebSocketSubprotocol},
		EnableCompression: true,
	})
	return f(ctx)
}
