package port

import vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/valueobject"

type TokenParser interface {
	ParseAccessToken(payload string) (vo.AccessTokenData, error)
	ParseRefreshToken(payload string) (vo.RefreshTokenData, error)
}
