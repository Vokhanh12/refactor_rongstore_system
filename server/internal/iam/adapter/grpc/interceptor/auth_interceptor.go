package grpc

import (
	"context"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/grpc"
	cmd "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/command"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/usecases"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthUnaryInterceptor(
	authenticateUC *usecases.AuthenticateUsecase,
) grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		md, _ := metadata.FromIncomingContext(ctx)

		values := md.Get("authorization")
		if len(values) == 0 {
			return handler(ctx, req)
		}

		result, err := authenticateUC.Execute(
			ctx,
			cmd.AuthenticateCommand{
				Token: values[0],
			},
		)

		if err != nil {
			return nil, core.ToGRPCError(
				err.Code,
				err.Message,
			)
		}

		ctx = ctxutil.WithUser(
			ctx,
			ctxutil.UserContext{
				UserID:      result.UserID,
				RoleKeyStrs: result.RoleKeyStrs,
			},
		)

		return handler(ctx, req)
	}
}
