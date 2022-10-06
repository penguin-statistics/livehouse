package main

import (
	"log"

	"google.golang.org/protobuf/proto"

	"github.com/penguin-statistics/livehouse/internal/model/pb"
)

func main() {
	req := pb.MatrixUpdateSubscribeReq{
		Id: &pb.MatrixUpdateSubscribeReq_StageId{
			StageId: 1,
		},
	}

	b, err := proto.Marshal(&req)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("update subscription with stageId = 1 (hex): %x", b)
}
