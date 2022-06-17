package controller

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/penguin-statistics/livehouse/internal/model/pb"
	"github.com/penguin-statistics/livehouse/internal/pkg/lhcore"
)

type IntakeDeps struct {
	fx.In

	DropSet *lhcore.DropSet
}

type Intake struct {
	pb.UnimplementedConnectedLiveServiceServer

	IntakeDeps
}

func RegisterIntake(serv *grpc.Server, deps IntakeDeps) {
	c := &Intake{
		IntakeDeps: deps,
	}
	pb.RegisterConnectedLiveServiceServer(serv, c)
}

func (c *Intake) PushReportBatch(ctx context.Context, req *pb.ReportBatchRequest) (*pb.ReportBatchACK, error) {
	log.Info().
		Interface("request", req).
		Msgf("received report batch")

	for _, report := range req.GetReport() {
		for _, drops := range report.GetDrops() {
			el := c.DropSet.GetOrCreateElement(report.GetStageId(), drops.GetItemId())
			el.Incr(1, drops.GetQuantity(), report.GetGeneration())
		}
	}

	log.Debug().
		Msgf("processed report batch")

	return &pb.ReportBatchACK{}, nil
}

func (c *Intake) PushMatrixBatch(ctx context.Context, req *pb.MatrixBatchRequest) (*pb.MatrixBatchACK, error) {
	log.Info().
		Interface("request", req).
		Msgf("received matrix batch")

	for _, matrix := range req.GetMatrix() {
		el := c.DropSet.GetOrCreateElement(matrix.GetStageId(), matrix.GetItemId())
		el.CutOut(matrix.GetTimes(), matrix.GetQuantity(), req.GetGeneration())
	}

	return &pb.MatrixBatchACK{
		Generation: req.Generation,
	}, nil
}

func (c *Intake) GetMatrixBatch(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	c.DropSet.CombineElements.Range(func(key, value any) bool {
		el := value.(*lhcore.DropElement)
		fmt.Printf("stage: %d, item: %d | times: %d, quantity: %d\n", el.StageID, el.ItemID, el.Times.Sum(), el.Quantity.Sum())
		return true
	})

	return &pb.Empty{}, nil
}
