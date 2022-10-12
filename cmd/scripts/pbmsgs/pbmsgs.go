package main

import (
	"log"

	"google.golang.org/protobuf/proto"

	"exusiai.dev/livehouse/internal/model/pb"
)

func main() {
	stageId := uint32(175)
	req := pb.MatrixUpdateSubscribeReq{
		Id: &pb.MatrixUpdateSubscribeReq_StageId{
			StageId: stageId,
		},
	}

	b, err := proto.Marshal(&req)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("update subscription with stageId = %d (hex): %x", stageId, b)
}
