package grpc

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	ucs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/jwt"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func AuthZUnaryInterceptor(
	authorize ucs.AuthorizeUsecase,
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
				_ = ctxutil.WithUser(ctx, ctxutil.UserContext{
					UserID: payload.UserID,
					Roles:  payload.Roles,
				})
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

		userctx, ok := ctxutil.User(ctx)
		if ok && userctx.UserID != "" && userctx.Roles != nil {

			allowed, err := authorize.Execute(ctx, command.AuthorizeCommand{
				UserID:     userctx.UserID,
				TenantID:   userctx.,
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
