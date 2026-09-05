package mapper

import (
	authv1rs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/auth/v1/resources"
	cmd "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/command"
)

func ToLoginCommand(req *authv1rs.LoginRequest) cmd.LoginCommand {
	return cmd.LoginCommand{
		Identifier: req.Identifier,
		Password:   req.Password,
	}
}

func FromLoginResult(result cmd.LoginCommandResult) *authv1rs.LoginResponse {
	return &authv1rs.LoginResponse{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}
}
