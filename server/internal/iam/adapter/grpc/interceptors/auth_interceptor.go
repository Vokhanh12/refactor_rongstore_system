package grpc

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/jwt"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		if values := md.Get("x-jwt-payload"); len(values) > 0 {
			if payload, err := jwt.DecodePayload(values[0]); err == nil {
				ctx = ctxutil.WithUser(ctx, ctxutil.UserContext{
					UserID:   payload.UserID,
					Roles:    payload.Roles,
					TenantID: payload.TenantID,
				})
			} else {
				// log decode failure
			}
		}
		return handler(ctx, req)
	}
}
