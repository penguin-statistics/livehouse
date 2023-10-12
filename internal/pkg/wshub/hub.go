package wshub

import (
	"sync"
	"time"

	ws "github.com/fasthttp/websocket"
	"github.com/gofiber/websocket/v2"
	"github.com/rs/zerolog/log"
)

type HubConfig struct {
	MaxMessageSize int64

	PingWait   time.Duration
	PingPeriod time.Duration
	PongWait   time.Duration

	WriteWait time.Duration
}

var DefaultHubConfig = HubConfig{
	MaxMessageSize: 1024 * 1024,

	PingWait:   10 * time.Second,
	PingPeriod: 1 * time.Minute,
	PongWait:   10 * time.Second,

	WriteWait: 10 * time.Second,
}

type Hub struct {
	// The map of all the clients, indexed by the client's ID.
	Clients sync.Map

	Config *HubConfig
}

func NewHub(config ...HubConfig) *Hub {
	h := &Hub{
		Clients: sync.Map{},
	}
	if len(config) > 0 {
		h.Config = &config[0]
	} else {
		h.Config = &DefaultHubConfig
	}
	return h
}

func (h *Hub) Add(c *Client) {
	h.Clients.Store(c.ID, c)
}

func (h *Hub) Remove(c *Client) {
	h.Clients.Delete(c.ID)
}

func (h *Hub) Broadcast(msg *ws.PreparedMessage) {
	h.Clients.Range(func(key, value any) bool {
		c := value.(*Client)
		c.Send <- msg
		return true
	})
}

func (h *Hub) NewClient(conn *websocket.Conn, id string) *Client {
	c := &Client{
		Conn: conn,
		Send: make(chan *ws.PreparedMessage, 8),
		Recv: make(chan []byte, 2),
		Done: make(chan struct{}),
		Hub:  h,
		ID:   id,
	}
	h.Add(c)
	return c
}

type Client struct {
	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan *ws.PreparedMessage

	// Buffered channel of inbound messages.
	Recv chan []byte

	// Channel to signal the termination of client
	Done chan struct{}

	// The hub.
	Hub *Hub

	// The client's ID.
	ID string

	destroyOnce sync.Once
}

func (c *Client) Spin() {
	go c.readPump()
	go c.writePump()
}

func (c *Client) readPump() {
	defer func() {
		c.Destroy()
	}()
	c.Conn.SetReadLimit(c.Hub.Config.MaxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(c.Hub.Config.PongWait + c.Hub.Config.PingPeriod))
	c.Conn.SetPongHandler(func(string) error {
		log.Debug().Msg("got pong from client")
		c.Conn.SetReadDeadline(time.Now().Add(c.Hub.Config.PongWait + c.Hub.Config.PingPeriod))
		return nil
	})
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}
		c.Recv <- message
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(c.Hub.Config.PingPeriod)
	defer func() {
		ticker.Stop()
		c.Destroy()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// The hub closed the channel.
				return
			}

			err := c.Conn.WritePreparedMessage(message)
			if err != nil {
				return
			}
		case <-ticker.C:
			log.Debug().Msg("sending ping to client")
			c.Conn.SetWriteDeadline(time.Now().Add(c.Hub.Config.WriteWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}

		case <-c.Done:
			return
		}
	}
}

func (c *Client) SendClose() error {
	return c.Conn.WriteControl(websocket.CloseMessage, nil, time.Now().Add(c.Hub.Config.WriteWait))
}

func (c *Client) Destroy() {
	c.destroyOnce.Do(func() {
		log.Info().Str("clientId", c.ID).Msg("closing client")
		c.Hub.Remove(c)
		c.Conn.Close()
		close(c.Done)
		close(c.Send)
		close(c.Recv)
	})
}
