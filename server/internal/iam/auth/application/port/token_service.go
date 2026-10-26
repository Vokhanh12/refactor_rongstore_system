package port

import (
	"os/user"

	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/valueobject"
)

type TokenParser interface {
	ParseAccessToken(payload string) (AccessTokenClaims, error)
}

type TokenSigner interface {
	SignAccessToken(
		userID string,
		roles []TokenRoleScope,
		authzVersion int,
	) (string, error)

	SignRefreshToken(
		userID string,
	) (string, error)
}

type TokenService interface {
	GenerateTokens(u *user.User) (*vo.TokenPair, error)

	ExtractUserIDFromAccessToken(token string) (string, error)
}
