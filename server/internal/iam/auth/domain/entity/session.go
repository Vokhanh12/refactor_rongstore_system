package entity

import "time"

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
	IsRevoked      bool
}

// Constructor đảm bảo khởi tạo đúng trạng thái
func NewSession(id, userID, deviceID, tenantID, ip, userAgent string, ttl time.Duration) *Session {
	now := time.Now()
	return &Session{
		Id:             id,
		UserID:         userID,
		DeviceID:       deviceID,
		TenantID:       tenantID,
		Ip:             ip,
		UserAgent:      userAgent,
		CreatedAt:      now,
		LastActivityAt: now,
		ExpiresAt:      now.Add(ttl),
		IsRevoked:      false,
	}
}

func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}

func (s *Session) Revoke() {
	s.IsRevoked = true
	now := time.Now()
	s.RevokedAt = &now
}

// Thêm hành vi cập nhật hoạt động cuối (dùng khi user gọi API định kỳ)
func (s *Session) Touch() {
	s.LastActivityAt = time.Now()
}
