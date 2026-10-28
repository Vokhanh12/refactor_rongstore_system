package port

import vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/valueobject"

type TokenSigner interface {
	SignAccessToken(
		userID string,
		roles []vo.TokenRoleScope,
		authzVersion int,
	) (string, error)

	SignRefreshToken(
		userID string,
	) (string, error)
}
