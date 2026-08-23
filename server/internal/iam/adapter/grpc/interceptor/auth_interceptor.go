package grpc

import (
	"context"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/grpc"
	cmd "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/command"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/usecases"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthUnaryInterceptor(
	authUc *usecases.AuthenticateUsecase,
) grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		values := md.Get("x-jwt-payload")
		if len(values) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing jwt payload")
		}

		result, err := authUc.Execute(
			ctx,
			cmd.AuthenticateCommand{
				Payload: values[0],
			},
		)

		if err != nil {
			return nil, core.ToGRPCError(
				err.Code,
				err.Message,
			)
		}

		return handler(ctx, req)
	}
}
