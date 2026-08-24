package security

import aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

type TokenParser interface {
	ParseAccessToken(payload string) (AccessTokenClaims, *aerrs.AppError)
}

type TokenSigner interface {
	SignAccessToken(
		userID string,
		roles []RoleScope,
		authzVersion int,
	) (string, *aerrs.AppError)

	SignRefreshToken(
		claims RefreshTokenClaims,
	) (string, *aerrs.AppError)
}
