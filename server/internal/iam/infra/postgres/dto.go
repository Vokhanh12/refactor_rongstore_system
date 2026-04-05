package postgres

type roleRefDTO struct {
	RoleCode string `json:"role_code"`
	ScopeID  string `json:"scope_id"`
}
