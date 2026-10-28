package usecase

import (
	"context"
	"time"

	// Dùng để sinh ID cho session (nếu id trong entity là string UUID)
	com "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/command"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/port"
	sec "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/security"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/entity" // <-- Import entity session
	wr "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/writer"
	errs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/errors"
	aerr "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type LoginLocalUsecase struct {
	passwordHasher      sec.PasswordHasher
	credentialWriter    wr.CredentialWriter
	tokenSigner         sec.TokenSigner
	authorizationReader port.AuthorizationReader
	sessionWriter       wr.SessionWriter
}

func NewLoginUsecase(
	ph sec.PasswordHasher,
	cw wr.CredentialWriter,
	ts sec.TokenSigner,
	ar port.AuthorizationReader,
	sw wr.SessionWriter,
) *LoginLocalUsecase {
	return &LoginLocalUsecase{
		passwordHasher:      ph,
		credentialWriter:    cw,
		tokenSigner:         ts,
		authorizationReader: ar,
		sessionWriter:       sw,
	}
}

func (u *LoginLocalUsecase) Execute(
	ctx context.Context,
	cmd com.LoginCommand,
) (com.LoginCommandResult, error) {

	identifier, err := vo.NewLoginIdentifier(cmd.Identifier)
	if err != nil {
		return com.LoginCommandResult{}, err
	}

	credential, err := u.credentialWriter.FindByIdentifier( // (Lưu ý: hình như đây là reader nhưng bạn đang đặt tên biến là credentialWriter, giữ nguyên theo code của bạn)
		ctx,
		identifier,
	)
	if err != nil {
		return com.LoginCommandResult{}, err
	}

	if !u.passwordHasher.Verify(
		cmd.Password,
		credential.PasswordHash,
	) {
		return com.LoginCommandResult{}, aerr.New(
			errs.INVALID_CREDENTIALS,
		)
	}

	roleScopes, err := u.authorizationReader.ListRoleScopes(
		ctx,
		credential.UserID,
	)
	if err != nil {
		return com.LoginCommandResult{}, err
	}

	// Gọi Constructor NewSession mà bạn đã thiết kế rất chuẩn ở tầng Entity
	newSession := entity.NewSession(
		credential.UserID.String(),
		cmd.DeviceID,
		cmd.TenantID,
		cmd.Ip,
		cmd.UserAgent,
	)

	// Lưu session xuống Database
	if err := u.sessionWriter.Create(ctx, newSession); err != nil {
		return com.LoginCommandResult{}, err
	}
	// =========================================================================

	// Tạo Access Token (chứa user ID, role scopes, auth version)
	accessToken, err := u.tokenSigner.SignAccessToken(
		credential.UserID.String(),
		roleScopes,
		credential.AuthzVersion,
	)
	if err != nil {
		return com.LoginCommandResult{}, err
	}

	// Tạo Refresh Token (Đưa cái sessionID vào để sau này khi user gọi /refresh,
	// hệ thống sẽ lấy sessionID này ra check xem session còn ACTIVE hay không)
	refreshToken, err := u.tokenSigner.SignRefreshToken(
		credential.UserID.String(),
		sessionID, // 👈 Gắn sessionID vào claims của refresh token (hoặc dùng chính sessionID làm opaque token)
	)
	if err != nil {
		return com.LoginCommandResult{}, err
	}

	return com.LoginCommandResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64((15 * time.Minute).Seconds()), // Đã fix cú pháp thời gian
	}, nil
}
