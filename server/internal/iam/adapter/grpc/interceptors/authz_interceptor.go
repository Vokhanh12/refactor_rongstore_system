package grpc

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	ucs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
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

		roleRefs := make([]vo.RoleRef, 0, len(userctx.Roles))

		for _, r := range userctx.Roles {
			roleRef, err := parseRoleRef(r)
			if err != nil {
				return nil, err
			}
			roleRefs = append(roleRefs, roleRef)
		}

		result, err := authorize.Execute(ctx, command.AuthorizeCommand{
			UserID:     userctx.UserID,
			RoleRef:    roleRefs,
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

func parseRoleRef(input string) (*vo.RoleRef, error) {
	parts := strings.SplitN(input, ":", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return nil, status.Errorf(codes.PermissionDenied, "invalid role format")
	}

	id, err := uuid.Parse(parts[1])
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "invalid role id")
	}

	return vo.NewRoleRef(id, parts[0])
}
