package writer

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/entity"
)

type SessionWriter interface {
	// Tạo mới session khi login/register
	Create(ctx context.Context, session *entity.Session) error

	// Tìm session theo ID (dùng khi check token / refresh token)
	FindByID(ctx context.Context, sessionID string) (*entity.Session, error)

	// Cập nhật trạng thái hoặc thông tin session (ví dụ: Revoke, Touch)
	Update(ctx context.Context, session *entity.Session) error

	// Thu hồi toàn bộ session của một user (dùng khi đổi mật khẩu toàn cục / đăng xuất tất cả thiết bị)
	RevokeAllByUserID(ctx context.Context, userID string) error
}
