package usecases

import (
	"context"

	com "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/command"
	jwtport "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/ports"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type AuthenticateUsecase struct {
	tokenDecoreder jwtport.TokenDecoreder
}

func NewAuthenticateUsecase(
	tokenDecoreder jwtport.TokenDecoreder,
) *AuthenticateUsecase {
	return &AuthenticateUsecase{
		tokenDecoreder: tokenDecoreder,
	}
}

func (u *AuthenticateUsecase) Execute(
	ctx context.Context,
	cmd com.AuthenticateCommand,
) (*com.AuthenticateCommandResult, *aerrs.AppError) {

	payload, err := u.tokenDecoreder.DecorePayload(cmd.Token)
	if err != nil {
		return nil, err
	}

	return &com.AuthenticateCommandResult{
		UserID:      payload.UserID,
		RoleKeyStrs: payload.RoleKeyStrs,
	}, nil
}
