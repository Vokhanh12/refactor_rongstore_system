package projection

import (
	"github.com/google/uuid"
	enu "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
)

type RoleView struct {
	ID          uuid.UUID
	Code        string
	Name        string
	Description *string

	ScopeType enu.RoleScopeType

	IsActive bool
	IsSystem bool
	IsSuper  bool

	Level int32
}
