package service

import (
	"context"
	"errors"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/entity"
	wr "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/writer"
)

type SessionService struct {
	sessionWr wr.SessionWriter
}

func NewSessionService(wr wr.SessionWriter) *SessionService {
	return &SessionService{
		sessionWr: wr,
	}
}

// ValidateAndGetSession là hàm dùng chung cho các Use Case để kiểm tra session nhanh chóng
func (s *SessionService) ValidateAndGetSession(ctx context.Context, sessionID string) (*entity.Session, error) {
	// 1. Lấy session từ database thông qua repository
	session, err := s.sessionWr.FindByID(ctx, sessionID)
	if err != nil {
		return nil, errors.New("session not found or database error")
	}

	// 2. Tận dụng hành vi IsValid() đã viết sẵn trong Entity để kiểm tra trạng thái và hạn sử dụng
	if !session.IsValid() {
		return nil, errors.New("session is expired, revoked, or invalid")
	}

	return session, nil
}
