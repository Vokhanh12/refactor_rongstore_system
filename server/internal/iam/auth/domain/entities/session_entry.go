package entities

import "time"

type SessionEntry struct {
	SessionID string
	ClientPub []byte
	ServerPub []byte

	Kc2s []byte
	Ks2c []byte

	HKDFSalt []byte
	Expiry   time.Time

	UserID   string
	ClientIP string
}
