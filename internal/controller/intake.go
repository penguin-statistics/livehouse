package controller

import (
	"context"

	"google.golang.org/grpc"

	"github.com/penguin-statistics/livehouse/internal/model/pb"
)

type IntakeController struct {
	pb.UnimplementedConnectedLiveServiceServer
}

func RegisterIntakeController(serv *grpc.Server) {
	c := &IntakeController{}
	pb.RegisterConnectedLiveServiceServer(serv, c)
}

func (c *IntakeController) PushReportBatch(ctx context.Context, req *pb.ReportBatchRequest) (*pb.ReportBatchACK, error) {
	return &pb.ReportBatchACK{
		Generation: req.Generation,
	}, nil
}

func (c *IntakeController) PushMatrixBatch(ctx context.Context, req *pb.MatrixBatchRequest) (*pb.MatrixBatchACK, error) {
	return &pb.MatrixBatchACK{
		Generation: req.Generation,
	}, nil
}
