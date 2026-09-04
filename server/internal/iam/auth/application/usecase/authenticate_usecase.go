package usecase

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/adapter/mapper"
	com "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/command"
	sec "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/security"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type AuthenticateUsecase struct {
	tokenParser sec.TokenParser
}

func NewAuthenticateUsecase(
	td sec.TokenParser,
) *AuthenticateUsecase {
	return &AuthenticateUsecase{
		tokenParser: td,
	}
}

func (u *AuthenticateUsecase) Execute(
	ctx context.Context,
	cmd com.AuthenticateCommand,
) (*com.AuthenticateCommandResult, *aerrs.AppError) {

	claims, err := u.tokenParser.ParseAccessToken(cmd.Payload)
	if err != nil {
		return &com.AuthenticateCommandResult{
			Allowed: false,
		}, err
	}

	identity := mapper.ToIdentityContext(claims)

	return &com.AuthenticateCommandResult{
		Allowed:  true,
		Identity: identity,
	}, nil
}
