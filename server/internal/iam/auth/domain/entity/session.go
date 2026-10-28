package entity

import (
	"time"
	"uuid"

	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/valueobject"
)

type Session struct {
	Id             string
	UserID         string
	DeviceID       string
	TenantID       string
	Ip             string
	UserAgent      string
	CreatedAt      time.Time
	LastActivityAt time.Time
	ExpiresAt      time.Time
	RevokedAt      *time.Time
	Status         vo.SessionStatus
}

// Constructor đảm bảo khởi tạo đúng trạng thái
func NewSession(userID, deviceID, tenantID, ip, userAgent string) *Session {
	now := time.Now()
	return &Session{
		Id:             uuid.New().String(),
		UserID:         userID,
		DeviceID:       deviceID,
		TenantID:       tenantID,
		Ip:             ip,
		UserAgent:      userAgent,
		CreatedAt:      now,
		LastActivityAt: now,
		ExpiresAt:      now.Add(7 * 24 * time.Hour),
		Status:         vo.StatusActive,
	}
}

// Kiểm tra session còn sống và hợp lệ không
func (s *Session) IsValid() bool {
	// 1. Dùng hàm IsValid() của chính Value Object Status để kiểm tra xem trạng thái có hợp lệ không
	if !s.Status.IsValid() || s.Status != vo.StatusActive {
		return false
	}
	// 2. Kiểm tra thời gian hết hạn
	if time.Now().After(s.ExpiresAt) {
		return false
	}
	return true
}

func (s *Session) Revoke() {
	s.Status = vo.StatusRevoked
	now := time.Now()
	s.RevokedAt = &now
}

func (s *Session) Touch() {
	s.LastActivityAt = time.Now()
}
