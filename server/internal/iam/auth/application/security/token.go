package security

import aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

type TokenParser interface {
	ParseAccessTokenClaims(payload string) (AccessTokenClaims, *aerrs.AppError)
}

type TokenSigner interface {
	SignAccessToken(
		userID string,
		tenantID string,
		roles []RoleScope,
		authzVersion int,
	) (string, *aerrs.AppError)

	SignRefreshToken(
		claims RefreshTokenClaims,
	) (string, *aerrs.AppError)
}
