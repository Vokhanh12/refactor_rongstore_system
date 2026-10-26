package valueobject

import "time"

type TokenPair struct {
	AccessToken  string
	RefreshToken string
	ExpiredAt    time.Time
}
