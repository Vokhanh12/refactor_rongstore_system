package entities

import "time"

type Session struct {
	session_id       string
	user_id          string
	device_id        string
	tenant_id        string
	ip               string
	user_agent       string
	created_at       time.Time
	last_activity_at time.Time
	expires_at       time.Time
	revoked_at       *time.Time
}
