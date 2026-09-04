package grpc

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	uc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecase"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func AuthZUnaryInterceptor(
	authorize uc.AuthorizeUsecase,
) grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		protoReq, ok := req.(proto.Message)
		if !ok {
			return handler(ctx, req)
		}

		authOpt, err := extractAuthOptions(protoReq)
		if err != nil || authOpt == nil {
			return handler(ctx, req)
		}

		identity, ok := ctxutil.Identity(ctx)
		if !ok {
			return nil, status.Error(
				codes.Internal,
				"identity context missing",
			)
		}

		if identity.UserID == "" {
			return nil, status.Error(
				codes.Internal,
				"invalid identity context",
			)
		}

		result, err := authorize.Execute(
			ctx,
			command.AuthorizeCommand{
				UserID:     identity.UserID,
				RoleScopes: identity.RoleScopes,
				Resource:   authOpt.Resource,
				Action:     authOpt.Action,
				ResourceID: extractResourceID(
					protoReq,
					authOpt.ResourceIDField,
				),
			},
		)

		if err != nil {
			return nil, err
		}

		if !result.Allowed {
			return nil, status.Error(
				codes.PermissionDenied,
				"unauthorized",
			)
		}

		return handler(ctx, req)
	}
}
