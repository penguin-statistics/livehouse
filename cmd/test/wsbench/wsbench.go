package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
	"google.golang.org/protobuf/proto"

	"github.com/penguin-statistics/livehouse/internal/model/pb"
)

func penguinV3GetRequest(path string) []byte {
	log.Println("Fetching", path)
	r := lo.Must(http.NewRequest("GET", "https://penguin-stats.io/api/v3alpha"+path, nil))
	r.Header.Set("Accept", "application/vnd.penguin.v3+json")
	resp := lo.Must(http.DefaultClient.Do(r))
	defer resp.Body.Close()
	return lo.Must(io.ReadAll(resp.Body))
}

func penguinItemIds() []uint32 {
	resp := penguinV3GetRequest("/items")
	gj := gjson.ParseBytes(resp).Array()
	ids := make([]uint32, len(gj))
	for i, v := range gj {
		ids[i] = uint32(v.Get("penguinItemId").Uint())
	}
	return ids
}

func penguinStageIds() []uint32 {
	resp := penguinV3GetRequest("/stages")
	gj := gjson.ParseBytes(resp).Array()
	ids := make([]uint32, len(gj))
	for i, v := range gj {
		ids[i] = uint32(v.Get("penguinStageId").Uint())
	}
	return ids
}

func randomInSlice(slice []uint32) uint32 {
	return slice[rand.Intn(len(slice))]
}

var (
	Servers = []pb.Server{pb.Server_CN, pb.Server_US, pb.Server_JP, pb.Server_KR}
	Items   = penguinItemIds()
	Stages  = penguinStageIds()
)

func getRandomSubscriptionReq() []byte {
	if rand.Intn(2) == 0 {
		return lo.Must(proto.Marshal(&pb.MatrixUpdateSubscribeReq{
			Id: &pb.MatrixUpdateSubscribeReq_StageId{
				StageId: randomInSlice(Stages),
			},
			Server: Servers[rand.Intn(len(Servers))],
		}))
	} else {
		return lo.Must(proto.Marshal(&pb.MatrixUpdateSubscribeReq{
			Id: &pb.MatrixUpdateSubscribeReq_ItemId{
				ItemId: randomInSlice(Items),
			},
			Server: Servers[rand.Intn(len(Servers))],
		}))
	}
}

func main() {
	for i := 0; i < 5000; i++ {
		conn, resp, err := websocket.DefaultDialer.Dial("ws://localhost:9020/live", http.Header{
			"Sec-WebSocket-Protocol": []string{"v3.penguin-stats.live+proto"},
		})
		if err != nil {
			log.Println("dial error:", err, resp)
			panic(err)
		}

		go func(i int) {
			err := conn.WriteMessage(websocket.BinaryMessage, getRandomSubscriptionReq())
			if err != nil {
				log.Println("write error:", err)
				return
			}

			log.Println("client", i, "is ready")

			for {
				_, _, err = conn.ReadMessage()
				if err != nil {
					log.Println("read error:", err)
				}

				// log.Println("client", i, "received", mt, ":", hex.EncodeToString(b))
			}
		}(i)

		time.Sleep(time.Millisecond * 1)
	}

	select {}
}
