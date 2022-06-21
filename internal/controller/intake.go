package controller

import (
	"context"

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
			el := c.DropSet.GetOrCreateElement(lhcore.IDSet{
				StageID: report.GetStageId(),
				ItemID:  drops.GetItemId(),
			})
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
		el := c.DropSet.GetOrCreateElement(lhcore.IDSet{
			StageID: matrix.GetStageId(),
			ItemID:  matrix.GetItemId(),
		})
		el.CutOut(matrix.GetTimes(), matrix.GetQuantity(), req.GetGeneration())
	}

	return &pb.MatrixBatchACK{
		Generation: req.Generation,
	}, nil
}
