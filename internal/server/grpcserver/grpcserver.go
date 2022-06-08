package grpcserver

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/penguin-statistics/livehouse/internal/config"
)

func Create(conf *config.Config) *grpc.Server {
	serv := grpc.NewServer()
	reflection.Register(serv)

	return serv

	// registerPromOnce.Do(func() {
	// 	fiberprom := fiberprometheus.New(observability.ServiceName)
	// 	fiberprom.RegisterAt(app, "/metrics")
	// })

	// if conf.TracingEnabled {
	// 	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint())
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	tracerProvider := tracesdk.NewTracerProvider(
	// 		tracesdk.WithSyncer(exporter),
	// 		tracesdk.WithResource(resource.NewWithAttributes(
	// 			semconv.SchemaURL,
	// 			semconv.ServiceNameKey.String("livehouse"),
	// 			attribute.String("environment", "dev"),
	// 		)),
	// 	)
	// 	otel.SetTracerProvider(tracerProvider)
	// }

	// if conf.DevMode {
	// 	log.Info().Msg("Running in DEV mode")
	// 	app.Use(pprof.New())
	// }

	// return app
}
