package grpcsvr

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/penguin-statistics/livehouse/internal/config"
)

func Create(conf *config.Config) *grpc.Server {
	serv := grpc.NewServer()
	reflection.Register(serv)

	return serv
}
