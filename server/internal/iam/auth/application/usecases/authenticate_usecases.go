package usecases

import (
	"context"

	com "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/command"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/security"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
)

type AuthenticateUsecase struct {
	tokenDecoreder security.TokenDecoreder
}

func NewAuthenticateUsecase(
	tokenDecoreder security.TokenDecoreder,
) *AuthenticateUsecase {
	return &AuthenticateUsecase{
		tokenDecoreder: tokenDecoreder,
	}
}

func (u *AuthenticateUsecase) Execute(
	ctx context.Context,
	cmd com.AuthenticateCommand,
) (*com.AuthenticateCommandResult, *aerrs.AppError) {

	token, err := u.tokenDecoreder.DecoreToken(cmd.Token)

	if err != nil {
		return &com.AuthenticateCommandResult{
			Allowed: false,
		}, err
	}

	ctx = ctxutil.WithUser(
		ctx,
		ctxutil.UserContext{
			UserID:      token.UserID,
			RoleKeyStrs: token.RoleKeyStrs,
		},
	)

	return &com.AuthenticateCommandResult{
		Allowed: true,
	}, nil
}
