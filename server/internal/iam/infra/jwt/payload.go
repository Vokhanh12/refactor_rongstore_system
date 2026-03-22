package jwt

type Payload struct {
	UserID string   `json:"sub"`
	OrgID  string   `json:"org,omitempty"`
	UnitID string   `json:"unit,omitempty"`
	Roles  []string `json:"roles,omitempty"`
}
