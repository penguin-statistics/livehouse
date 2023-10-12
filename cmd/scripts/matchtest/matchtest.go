package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fasthttp/websocket"
	"github.com/samber/lo"
	"google.golang.org/protobuf/proto"

	"exusiai.dev/gommon/model/pb"
)

var (
	StageID    = uint32(18)
	Generation = uint64(0)
)

func genreq(drops []*pb.Drop) *pb.ReportBatchRequest {
	return &pb.ReportBatchRequest{
		Reports: []*pb.Report{
			{
				Server:     pb.Server_CN,
				Generation: Generation,
				StageId:    StageID,
				Drops:      drops,
			},
		},
	}
}

func main() {
	// conn := lo.Must(grpc.Dial("localhost:9015", grpc.WithInsecure()))

	// client := pb.NewConnectedLiveServiceClient(conn)
	// lo.Must(client.PushReportBatch(context.Background(), genreq([]*pb.Drop{{
	// 	ItemId:   1,
	// 	Quantity: 1,
	// }})))

	req := pb.MatrixUpdateSubscribeReq{
		Id: &pb.MatrixUpdateSubscribeReq_ItemId{
			ItemId: 38,
		},
	}

	b, err := proto.Marshal(&req)
	if err != nil {
		log.Fatalln(err)
	}

	wsconn, resp, err := websocket.DefaultDialer.Dial("ws://localhost:9020/live", http.Header{
		"Sec-WebSocket-Protocol": []string{"v3.penguin-stats.live+proto"},
	})
	if err != nil {
		log.Fatalln("dial error:", err, resp)
	}

	err = wsconn.WriteMessage(websocket.BinaryMessage, b)
	if err != nil {
		log.Fatalln("write error:", err)
	}

	go func() {
		for {
			_, msg, err := wsconn.ReadMessage()
			if err != nil {
				log.Fatalln("read error:", err)
			}

			var msgresp pb.MatrixUpdateMessage
			proto.Unmarshal(msg, &msgresp)

			log.Print("received the following updates:\n", lo.Reduce(msgresp.Segments, func(acc string, seg *pb.MatrixUpdateMessage_Element, _ int) string {
				return acc + fmt.Sprintf("  - server=%s, stage=%d, item=%d, times=%d, quantity=%d\n", seg.Server, seg.StageId, seg.ItemId, seg.Times, seg.Quantity)
			}, ""))
		}
	}()

	// lo.Must(client.PushReportBatch(context.Background(), genreq([]*pb.Drop{})))

	// log.Println("expect quantity = 1, times = 2")
	// time.Sleep(time.Second * 10)

	// lo.Must(client.PushReportBatch(context.Background(), genreq([]*pb.Drop{
	// 	{
	// 		ItemId:   1,
	// 		Quantity: 1,
	// 	},
	// })))

	// log.Println("expect quantity = 2, times = 3")
	// time.Sleep(time.Second * 10)

	// lo.Must(client.PushReportBatch(context.Background(), genreq([]*pb.Drop{
	// 	{
	// 		ItemId:   1,
	// 		Quantity: 2,
	// 	},
	// })))

	// log.Println("expect quantity = 4, times = 4")
	// time.Sleep(time.Second * 10)

	// lo.Must(client.PushReportBatch(context.Background(), genreq([]*pb.Drop{})))

	// log.Println("expect quantity = 4, times = 5")

	select {}
}
