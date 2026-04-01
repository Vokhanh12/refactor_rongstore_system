package grpc

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	ucs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func AuthZUnaryInterceptor(authorize ucs.AuthorizeUsecase) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		protoReq, ok := req.(proto.Message)
		if !ok {
			return handler(ctx, req)
		}

		authOpt, err := extractAuthOptions(protoReq)
		if err != nil || authOpt == nil {
			return handler(ctx, req)
		}

		userctx, ok := ctxutil.User(ctx)
		if !ok || userctx.UserID == "" || userctx.Roles == nil {
			return nil, status.Errorf(codes.PermissionDenied, "unauthenticated")
		}

		result, err := authorize.Execute(ctx, command.AuthorizeCommand{
			UserID:     userctx.UserID,
			Roles:      userctx.Roles,
			Resource:   authOpt.Resource,
			Action:     authOpt.Action,
			ResourceID: extractResourceID(protoReq, authOpt.ResourceIDField),
		})

		if err != nil {
			return nil, err
		}

		if !result.Allowed {
			return nil, status.Errorf(codes.PermissionDenied, "unauthorized")
		}

		return handler(ctx, req)
	}
}
