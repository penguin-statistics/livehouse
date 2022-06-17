package controller

import (
	"time"

	ws "github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"github.com/penguin-statistics/livehouse/internal/constant"
	"github.com/penguin-statistics/livehouse/internal/service"
	"github.com/penguin-statistics/livehouse/internal/util"
)

var InvalidSubprotocolMsg = util.Must(ws.NewPreparedMessage(websocket.CloseMessage, []byte("invalid subprotocol: expect subprotocol "+constant.LiveWebSocketSubprotocol)))

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
			conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
			err := conn.WritePreparedMessage(InvalidSubprotocolMsg)
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
