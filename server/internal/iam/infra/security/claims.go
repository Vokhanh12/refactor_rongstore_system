package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/valueobject"
)

type accessTokenClaims struct {
	vo.AccessTokenData `json:"data"`
	jwt.RegisteredClaims
}

type refreshTokenClaims struct {
	vo.RefreshTokenData `json:"data"`
	jwt.RegisteredClaims
}
