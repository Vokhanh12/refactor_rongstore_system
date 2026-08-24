package usecases

import (
	"context"

	com "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/command"
	sec "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/security"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type LoginUsecase struct {
	tokenParser sec.TokenParser
}

func NewLoginUsecase(
	td sec.TokenParser,
) *LoginUsecase {
	return &LoginUsecase{
		tokenParser: td,
	}
}

func (u *LoginUsecase) Execute(
	ctx context.Context,
	cmd com.LoginCommand,
) (*com.LoginCommandResult, *aerrs.AppError) {

	return &com.LoginCommandResult{
		Allowed: true,
	}, nil
}
