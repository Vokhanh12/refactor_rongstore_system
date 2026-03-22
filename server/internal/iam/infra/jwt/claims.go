package jwt

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID string   `json:"sub"`
	OrgID  string   `json:"org,omitempty"`
	UnitID string   `json:"unit,omitempty"`
	Roles  []string `json:"roles,omitempty"`

	jwt.RegisteredClaims
}
