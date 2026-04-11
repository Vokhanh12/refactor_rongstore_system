package valueobjects

import (
	"github.com/google/uuid"
	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type RoleRef struct {
	roleCode string
	scopeID  *uuid.UUID
}

func NewRoleRef(scopeID *uuid.UUID, roleCode string) *RoleRef {
	return &RoleRef{
		roleCode: roleCode,
		scopeID:  scopeID,
	}
}

func (p *RoleRef) validate() []aerrs.AppErrorDetail {

	var details []aerrs.AppErrorDetail

	if p.roleCode == "" {
		details = append(details, *aerrs.NewDetail(
			domain.REASON_REQUIRED,
			aerrs.WithField("roleCode"),
			aerrs.WithMessageDetail("role code is required"),
		))
		return details
	}

	return details
}

func (r RoleRef) RoleCode() string    { return r.roleCode }
func (r RoleRef) ScopeID() *uuid.UUID { return r.scopeID }

func (r RoleRef) String() string {
	return r.roleCode + ":" + r.scopeID.String()
}

type NewRoleRefParms struct {
	RoleCode string
	ScopeID  *uuid.UUID
}

func NewRoleRefFromPersistence(p NewRoleRefParms) RoleRef {
	return RoleRef{
		roleCode: p.RoleCode,
		scopeID:  p.ScopeID,
	}
}
