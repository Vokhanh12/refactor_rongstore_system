package grpc

import (
	"context"

	ucs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func AuthZUnaryInterceptor(
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
