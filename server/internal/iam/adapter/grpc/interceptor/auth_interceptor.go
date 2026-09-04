package grpc

import (
	"context"

	cmd "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/command"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/usecase"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthUnaryInterceptor(
	authUsecase *usecase.AuthenticateUsecase,
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

		result, err := authUsecase.Execute(
			ctx,
			cmd.AuthenticateCommand{
				Payload: values[0],
			},
		)

		if err != nil {
			return nil, err
		}

		ctx = ctxutil.WithIdentity(ctx, result.Identity)

		return handler(ctx, req)
	}
}
