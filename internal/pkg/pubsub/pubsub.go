package pubsub

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Message any

type Config struct {
	// batchinterv is the minimum time between messages.
	// It is guaranteed that the subscriber will receive
	// at most 1 batch of messages per batch interval.
	batchinterv time.Duration

	// pubtimeout is the maximum wait time for each
	// message to be published. If a message is not published
	// within the timeout, it will be dropped to the corresponding
	// subscriber, but may still be delivered to other subscribers.
	pubtimeout time.Duration
}

func NewConfig(batchinterv time.Duration, pubtimeout time.Duration) Config {
	return Config{
		batchinterv: batchinterv,
		pubtimeout:  pubtimeout,
	}
}

// PubSub is a simple pubsub implementation, with message
// throttling and batching support. It is designed to be used
// with the fiber/websocket package.
// A zero value is ready to use.
type PubSub struct {
	Config

	buf chan *Message
}

func NewPubSub(config Config) *PubSub {
	return &PubSub{
		Config: config,
		buf:    make(chan *Message, 64),
	}
}

type Subscriber struct {
	// id is the identifier of the subscriber. It is mainly
	// used for debugging purposes and thus requires to be
	// a fmt.GoStringer.
	id fmt.GoStringer

	// buf is the buffer of messages to be sent to the subscriber.
	buf []*Message

	// out is an unbuffered channel to send the batch of messages to the subscriber.
	out chan []*Message

	// done is an unbuffered channel to signal the fanner to stop.
	done chan struct{}

	Config
}

func NewSubscriber(id fmt.GoStringer, config Config) *Subscriber {
	s := &Subscriber{
		id:     id,
		buf:    make([]*Message, 0, 128),
		out:    make(chan []*Message),
		done:   make(chan struct{}),
		Config: config,
	}
	go s.fanner()
	return s
}

func (s *Subscriber) fanner() {
	t := time.NewTimer(s.batchinterv)

	for {
		select {
		case <-s.done:
			log.Trace().Func(func(e *zerolog.Event) {
				e.Str("id", s.id.GoString()).Msg("fanner: received done")
			})
			return
		case <-t.C:
			log.Trace().Func(func(e *zerolog.Event) {
				e.Str("id", s.id.GoString()).Msg("fanner: flushing")
			})
			s.flush()
			// reset after flush: don't let the timer fire before the last
			// flush has been sent.
			t.Reset(s.batchinterv)
		}
	}
}

func (s *Subscriber) flush() {
	if len(s.buf) == 0 {
		log.Trace().Func(func(e *zerolog.Event) {
			e.Str("id", s.id.GoString()).Msg("flush: no messages. skipping")
		})
		return
	}
	select {
	case <-s.done:
		log.Trace().Func(func(e *zerolog.Event) {
			e.Str("id", s.id.GoString()).Msg("flush: received done")
		})
		return
	case <-time.After(s.pubtimeout):
		log.Trace().Func(func(e *zerolog.Event) {
			e.Str("id", s.id.GoString()).Msg("flush: timed out")
		})
		return
	case s.out <- s.buf:
		log.Trace().Func(func(e *zerolog.Event) {
			e.Str("id", s.id.GoString()).Msg("flush: sent")
		})
		return
	}
}

func (s *Subscriber) ID() any {
	return s.id
}

func (s *Subscriber) Out() <-chan []*Message {
	return s.out
}

func (s *Subscriber) Recv(msg *Message) {
	if cap(s.buf) <= len(s.buf) {
		log.Trace().Func(func(e *zerolog.Event) {
			e.Str("id", s.id.GoString()).Msg("recv: buffer full")
		})
		return
	}
	s.buf = append(s.buf, msg)
}

func (s *Subscriber) Done() <-chan struct{} {
	return s.done
}
