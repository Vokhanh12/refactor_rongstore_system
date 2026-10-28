package mapper

import (
	"context"

	authv1rs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/auth/v1/resources"
	cmd "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/command"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
)

func ToLoginCommand(ctx context.Context, req *authv1rs.LoginRequest) cmd.LoginCommand {
	reqCtx := ctxutil.MustRequest(ctx)
	return cmd.LoginCommand{
		Identifier: req.Data.Identifier,
		Password:   req.Data.Password,
		DeviceID:   req.Data.DeviceId,
		TenantID:   reqCtx.TenantID,
		Ip:         reqCtx.ClientIP,
		UserAgent:  reqCtx.UserAgent,
	}
}

func FromLoginResult(result cmd.LoginCommandResult) *authv1rs.LoginResponse {
	return &authv1rs.LoginResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}
}
