package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/samber/lo"
	"github.com/tidwall/gjson"
	"google.golang.org/grpc"

	"exusiai.dev/gommon/model/pb"
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

var Servers = []pb.Server{pb.Server_CN, pb.Server_US, pb.Server_JP, pb.Server_KR}

func main() {
	conn, err := grpc.Dial("localhost:9015", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	rand.Seed(time.Now().UnixNano())

	cl := pb.NewConnectedLiveServiceClient(conn)

	items := penguinItemIds()
	stages := penguinStageIds()

	for {
		fakeReportSize := rand.Intn(100) + 200

		start := time.Now()
		var req pb.ReportBatchRequest
		req.Reports = make([]*pb.Report, fakeReportSize)
		for i := 0; i < fakeReportSize; i++ {
			drops := make([]*pb.Drop, rand.Intn(10)+1)
			for j := range drops {
				drops[j] = &pb.Drop{
					ItemId:   randomInSlice(items),
					Quantity: rand.Uint64()%50 + 10,
				}
			}

			req.Reports[i] = &pb.Report{
				Server:  Servers[rand.Intn(len(Servers))],
				StageId: randomInSlice(stages),
				Drops:   drops,
			}
		}
		log.Println("generated", fakeReportSize, "reports, took", time.Since(start))
		start = time.Now()
		_, err := cl.PushReportBatch(context.Background(), &req)
		if err != nil {
			log.Println("    failed to send report:", err)
		}
		log.Printf("sent reports and got ack. took %s", time.Since(start))
		time.Sleep(500 * time.Millisecond)
	}
}
