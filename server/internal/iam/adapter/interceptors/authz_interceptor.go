package grpc

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/domain/ports"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/jwt"

	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/util/ctxutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func AuthZUnaryInterceptor(
	store ports.IRedisCache,
	authz ports.IAuthzUsecase,
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

		protoReq, ok := req.(proto.Message)
		if !ok {
			return handler(ctx, req)
		}

		authOpt, err := extractAuthOptions(protoReq)
		if err != nil || authOpt == nil {
			return handler(ctx, req)
		}

		resourceID := ""
		if authOpt.ResourceIDField != "" {
			resourceID, _ = extractResourceID(protoReq, authOpt.ResourceIDField)
		}

		if ctxutil.UserIDFromContext != nil && ctxutil.TenantIDFromContext != nil && ctxutil.RolesFromContext != nil {

			err := authz.Authorize(ctx, ports.AuthorizeInput{
				UserID:     ctxutil.UserIDFromContext,
				TenantID:   ctxutil.TenantIDFromContext,
				Roles:      ctxutil.RolesFromContext,
				Resource:   authOpt.Resource,
				Action:     authOpt.Action,
				ResourceID: resourceID,
			})

			if err != nil {
				return nil, err
			}
		}

		return handler(ctx, req)
	}
}
