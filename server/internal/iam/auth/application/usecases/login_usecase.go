package usecases

import (
	"context"
	"time"

	com "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/command"
	sec "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/security"
	repo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/repository"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/valueobject"
	errs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/errors"
	aerr "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type LoginUsecase struct {
	passwordHasher sec.PasswordHasher
	credentialRepo repo.CredentialRepository
	tokenSigner    sec.TokenSigner
}

func NewLoginUsecase(
	passwordHasher sec.PasswordHasher,
	credentialRepo repo.CredentialRepository,
	tokenSigner sec.TokenSigner,
) *LoginUsecase {
	return &LoginUsecase{
		passwordHasher: passwordHasher,
		credentialRepo: credentialRepo,
		tokenSigner:    tokenSigner,
	}
}

func (u *LoginUsecase) Execute(
	ctx context.Context,
	cmd com.LoginCommand,
) (*com.LoginCommandResult, error) {

	identifier, err := vo.NewLoginIdentifier(cmd.Identifier)
	if err != nil {
		return nil, err
	}

	credential, err := u.credentialRepo.FindByIdentifier(
		ctx,
		identifier.Value(),
	)

	if err != nil {
		return nil, err
	}

	if !u.passwordHasher.Verify(
		cmd.Password,
		credential.PasswordHash,
	) {
		return nil, aerr.New(
			errs.INVALID_CREDENTIALS,
		)
	}

	accessToken, err := u.tokenSigner.SignAccessToken(
		credential.UserID.String(),
		nil,
		credential.AuthzVersion,
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := u.tokenSigner.SignRefreshToken(
		sec.RefreshTokenClaims{
			UserID:    credential.UserID.String(),
			ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
		},
	)

	if err != nil {
		return nil, err
	}

	return &com.LoginCommandResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64((15 * time.Minute).Seconds()),
	}, nil
}
