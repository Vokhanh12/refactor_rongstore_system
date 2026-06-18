package main

import (
	"context"
	"log"
	"net"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	iampb "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/v1/services"
	wire "github.com/vokhanh12/refactor-rongstore-system/server/internal"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/config"

	obs_grpc "github.com/vokhanh12/refactor-rongstore-system/server/pkg/observability/grpc"
)

func main() {
	ctx := context.Background()
	cfg := config.Load()

	if err := logger.Init(
		logger.WithService("iam-service"),
	); err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}

	shutdownTracer, err := trace.Init(ctx, *cfg)
	if err != nil {
		log.Fatalf("failed to init tracer: %v", err)
	}
	defer shutdownTracer(ctx)

	deps := wire.InitializeIamHandler(ctx, cfg)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(
			otelgrpc.NewServerHandler(
				otelgrpc.WithTracerProvider(otel.GetTracerProvider()),
			),
		),
		grpc.ChainUnaryInterceptor(

			obs_grpc.RequestContextInterceptor(),

			obs_grpc.LoggingUnaryInterceptor("iam-service"),
			obs_grpc.MetricsUnaryInterceptor("iam-service"),

			obs_grpc.AuthUnaryInterceptor(
				deps.RedisSessionStore,
				auth.DefaultGrpcRules(),
			),
		),
	)

	reflection.Register(grpcServer)
	iampb.RegisterIamServiceServer(grpcServer, deps.Handler)

	log.Print("gRPC server started on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpc server stopped: %v", err)
	}
}
