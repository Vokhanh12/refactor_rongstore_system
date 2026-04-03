package valueobjects

import (
	"strings"

	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type RoleRef struct {
	roleCode string
	scopeID  string
}

func NewRoleRef(value string) (*RoleRef, []aerrs.AppErrorDetail) {
	var details []aerrs.AppErrorDetail

	if value == "" {
		details = append(details, *aerrs.NewDetail(
			domain.REASON_REQUIRED,
			aerrs.WithField("RoleRef"),
			aerrs.WithMessageDetail("RoleRef is required"),
		))
		return nil, details
	}

	parts := strings.Split(value, ":")
	if len(parts) != 2 {
		details = append(details, *aerrs.NewDetail(
			domain.REASON_INVALID_FORMAT,
			aerrs.WithField("role"),
			aerrs.WithMessageDetail("role must be in format ROLE:SCOPE"),
		))
		return nil, details
	}

	roleCode := strings.TrimSpace(parts[0])
	scopeID := strings.TrimSpace(parts[1])

	if roleCode == "" || scopeID == "" {
		details = append(details, *aerrs.NewDetail(
			domain.REASON_INVALID_FORMAT,
			aerrs.WithField("RoleRef"),
			aerrs.WithMessageDetail("role or scope is empty"),
		))
		return nil, details
	}

	return &RoleRef{
		roleCode: roleCode,
		scopeID:  scopeID,
	}, nil
}

func (r RoleRef) RoleCode() string { return r.roleCode }
func (r RoleRef) ScopeID() string  { return r.scopeID }

func (r RoleRef) Key() string {
	return r.roleCode + ":" + r.scopeID
}
