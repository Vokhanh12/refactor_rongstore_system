package valueobject

type SessionStatus string

const (
	StatusActive  SessionStatus = "ACTIVE"
	StatusRevoked SessionStatus = "REVOKED"
	StatusExpired SessionStatus = "EXPIRED"
)

func (t SessionStatus) IsValid() bool {
	switch t {
	case StatusActive, StatusRevoked, StatusExpired:
		return true
	}
	return false
}
