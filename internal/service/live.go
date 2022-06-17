package service

import (
	ws "github.com/fasthttp/websocket"
	"github.com/gofiber/websocket/v2"
	"github.com/rs/zerolog/log"

	"github.com/penguin-statistics/livehouse/internal/pkg/lhcore"
	"github.com/penguin-statistics/livehouse/internal/pkg/wshub"
)

type Live struct {
	DropSet *lhcore.DropSet
	Hub     *wshub.Hub
}

func NewLive(dropSet *lhcore.DropSet, hub *wshub.Hub) *Live {
	return &Live{
		DropSet: dropSet,
		Hub:     hub,
	}
}

func (l *Live) Handle(c *websocket.Conn) {
	log.Info().Msg("new live connection")

	cl := l.Hub.NewClient(c)
	cl.Spin()

	go func() {
		for {
			msg, ok := <-cl.Recv
			if !ok {
				log.Trace().Msg("client disconnected")
				return
			}

			log.Trace().Interface("msg", msg).Msg("received message")
			prepared, _ := ws.NewPreparedMessage(ws.BinaryMessage, msg)
			cl.Send <- prepared
		}
	}()

	go func() {
		for {
			prepared, _ := ws.NewPreparedMessage(ws.TextMessage, []byte("hello"))
			cl.Send <- prepared
		}
	}()

	<-cl.Done
}
