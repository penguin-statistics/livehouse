package controller

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/penguin-statistics/livehouse/internal/model/pb"
	"github.com/penguin-statistics/livehouse/internal/pkg/lhcore"
	"github.com/penguin-statistics/livehouse/internal/pkg/pgconv"
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
	log.Trace().
		Str("request", lo.Reduce(req.Reports, func(acc string, r *pb.Report, _ int) string {
			return acc + fmt.Sprintf("%s %d %v; ", r.GetServer(), r.GetStageId(), r.GetDrops())
		}, "")).
		Msgf("received report batch")

	for _, report := range req.GetReports() {
		server := pgconv.ServerIDFPBE(report.Server)
		idset := lhcore.IDSet{
			ServerID: server,
			StageID:  report.StageId,
		}

		for _, drops := range report.GetDrops() {
			idset.ItemID = drops.ItemId

			el := c.DropSet.GetOrCreateElement(idset)
			// not Incr-ing the times here.
			el.Incr(0, drops.Quantity, report.Generation)
		}

		c.DropSet.IncrTimes(report.StageId, server, report.Generation)
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
			ServerID: pgconv.ServerIDFPBE(req.Server),
			StageID:  matrix.GetStageId(),
			ItemID:   matrix.GetItemId(),
		})
		el.CutOut(matrix.GetTimes(), matrix.GetQuantity(), req.GetGeneration())
	}

	return &pb.MatrixBatchACK{
		Generation: req.Generation,
	}, nil
}
