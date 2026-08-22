package security

import aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

type TokenDecoder interface {
	DecodeAccessToken(
		encoded string,
	) (*AccessTokenClaims, *aerrs.AppError)

	DecodeRefreshToken(
		encoded string,
	) (*RefreshTokenClaims, *aerrs.AppError)
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
