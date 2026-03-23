package grpc

import (
	"context"

	ucs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthUnaryInterceptor(
	authzUc ucs.AuthorizeUsecase,
) grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		md, _ := metadata.FromIncomingContext(ctx)

		values := md.Get("x-jwt-payload")
		if len(values) > 0 {

			payload, err := jwt.DecodePayload(values[0])
			if err == nil {

				userCtx := auth.ToUserContext(payload)
				ctx = auth.WithUserContext(ctx, userCtx)
			}
		}
		return handler(ctx, req)
	}
}
