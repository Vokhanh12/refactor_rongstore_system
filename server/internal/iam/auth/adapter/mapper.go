package adapter

import (
	"context"

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
