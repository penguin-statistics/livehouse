package service

import (
	"encoding/base64"
	"time"

	ws "github.com/fasthttp/websocket"
	"github.com/gofiber/websocket/v2"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"

	"exusiai.dev/livehouse/internal/model/pb"
	"exusiai.dev/livehouse/internal/pkg/lhcore"
	"exusiai.dev/livehouse/internal/pkg/pgconv"
	"exusiai.dev/livehouse/internal/pkg/wshub"
)

type Live struct {
	DropSet *lhcore.DropSet
	Hub     *wshub.Hub

	ResultFlushInterval time.Duration
}

func NewLive(dropSet *lhcore.DropSet, hub *wshub.Hub) *Live {
	return &Live{
		DropSet:             dropSet,
		Hub:                 hub,
		ResultFlushInterval: 2 * time.Second,
	}
}

func (l *Live) Handle(c *websocket.Conn) {
	log.Info().Msg("new live connection")

	// {
	// 	item, err := proto.Marshal(&pb.MatrixUpdateSubscribeReq{
	// 		Header: &pb.Header{
	// 			Type: pb.MessageType_MATRIX_UPDATE_SUBSCRIBE_REQ,
	// 		},
	// 		Id: &pb.MatrixUpdateSubscribeReq_ItemId{
	// 			ItemId: 1,
	// 		},
	// 	})
	// 	log.Debug().
	// 		Err(err).
	// 		Str("base64", base64.StdEncoding.EncodeToString(item)).
	// 		Msg("example proto message: update subscription with itemId = 1")

	// 	stage, err := proto.Marshal(&pb.MatrixUpdateSubscribeReq{
	// 		Header: &pb.Header{
	// 			Type: pb.MessageType_MATRIX_UPDATE_SUBSCRIBE_REQ,
	// 		},
	// 		Id: &pb.MatrixUpdateSubscribeReq_StageId{
	// 			StageId: 1,
	// 		},
	// 	})
	// 	log.Debug().
	// 		Err(err).
	// 		Str("base64", base64.StdEncoding.EncodeToString(stage)).
	// 		Msg("example proto message: update subscription with stageId = 1")
	// }

	id := c.Locals("requestid").(string)
	log.Info().Str("clientId", id).Msg("connected")

	cl := l.Hub.NewClient(c, id)
	cl.Spin()
	defer cl.Destroy()

	sub := lhcore.NewSub(id)
	defer l.DropSet.RemoveSub(sub)

	go func() {
		for {
			msg, ok := <-cl.Recv
			if !ok {
				log.Debug().Str("clientId", id).Msg("client disconnected (service read)")
				return
			}

			log.Trace().Interface("msg", msg).Msg("received message")

			var req pb.MatrixUpdateSubscribeReq
			err := proto.Unmarshal(msg, &req)
			if err != nil {
				log.Error().Err(err).Msg("failed to unmarshal message")
				continue
			}

			var resp pb.MatrixUpdateSubscribeResp
			resp.Header = &pb.Header{
				Type: pb.MessageType_MATRIX_UPDATE_SUBSCRIBE_RESP,
			}

			switch {
			case req.GetItemId() != 0:
				log.Debug().
					Uint32("itemId", req.GetItemId()).
					Msg("replace subscription to item elements")

				err := l.DropSet.ReplaceSubToItemElements(req.GetItemId(), pgconv.ServerIDFPBE(req.Server), sub)
				if err != nil {
					log.Error().Err(err).Msg("failed to replace subscription to item elements")
					resp.Error = "failed to replace subscription to item elements: " + err.Error()
					continue
				}
			case req.GetStageId() != 0:
				log.Debug().
					Uint32("stageId", req.GetStageId()).
					Msg("replace subscription to stage elements")

				err := l.DropSet.ReplaceSubToStageElements(req.GetStageId(), pgconv.ServerIDFPBE(req.Server), sub)
				if err != nil {
					log.Error().Err(err).Msg("failed to replace subscription to stage elements")
					resp.Error = "failed to replace subscription to stage elements: " + err.Error()
				}
			default:
				log.Warn().Msg("failed to determine update subscription request: both stageId & itemId are zero-value")
				resp.Error = "failed to determine update subscription request: both stageId & itemId are zero-value"
			}
		}
	}()

	go func() {
		// result receiver & websocket sender
		timer := time.NewTicker(l.ResultFlushInterval)
		for {
			select {
			case <-timer.C:
				log.Debug().Str("clientId", cl.ID).Msg("flushing updates")

				elements := sub.Flush()
				if len(elements) == 0 {
					log.Debug().Str("clientId", cl.ID).Msg("no updates to flush")
					continue
				}

				msgels := make([]*pb.MatrixUpdateMessage_Element, 0, len(elements))
				for _, e := range elements {
					msgels = append(msgels, &pb.MatrixUpdateMessage_Element{
						Server:   pgconv.ServerIDTPBE(e.ServerID),
						StageId:  e.StageID,
						ItemId:   e.ItemID,
						Times:    e.Times,
						Quantity: e.Quantity,
					})
				}
				log.Debug().
					Interface("elements", msgels).
					Msg("elements to send")

				msg, err := proto.Marshal(&pb.MatrixUpdateMessage{
					Header: &pb.Header{
						Type: pb.MessageType_MATRIX_UPDATE_MESSAGE,
					},
					Segments: msgels,
				})
				if err != nil {
					log.Error().Err(err).Msg("failed to marshal message")
					continue
				}

				log.Debug().
					Str("base64", base64.StdEncoding.EncodeToString(msg)).
					Msg("sending update message")

				wsmsg, err := ws.NewPreparedMessage(ws.BinaryMessage, msg)
				if err != nil {
					log.Error().Err(err).Msg("failed to prepare websocket message")
					continue
				}

				cl.Send <- wsmsg
			case <-cl.Done:
				log.Debug().Str("clientId", id).Msg("client disconnected (service write)")
				return
			}
		}
	}()

	<-cl.Done
}
