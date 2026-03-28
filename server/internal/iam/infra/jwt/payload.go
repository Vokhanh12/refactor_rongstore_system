package jwt

type Payload struct {
	UserID string   `json:"sub"`
	Roles  []string `json:"roles,omitempty"`
}
