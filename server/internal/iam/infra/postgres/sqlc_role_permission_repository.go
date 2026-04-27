package postgres

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	en "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ re.RolePermissionRepository = (*SqlcRolePermissionRepository)(nil)

type SqlcRolePermissionRepository struct {
	queries *db.Queries
	dberr   dberr.DBError
}

func NewSqlcRolePermissionRepository(queries *db.Queries) repositories.RolePermissionRepository {
	return &SqlcRolePermissionRepository{queries: queries}
}

// Create implements [repositories.RolePermissionRepository].
func (s *SqlcRolePermissionRepository) Create(ctx context.Context, rolePermission *en.RolePermission) (*en.RolePermission, *apperrors.AppError) {
	panic("unimplemented")
}

// Delete implements [repositories.RolePermissionRepository].
func (s *SqlcRolePermissionRepository) Delete(ctx context.Context, id uuid.UUID) *apperrors.AppError {
	panic("unimplemented")
}

func (s *SqlcRolePermissionRepository) FindAllByRoleRefs(
	ctx context.Context,
	roleRefs []valueobjects.RoleRef,
) ([]*en.RolePermission, *apperrors.AppError) {

	if len(roleRefs) == 0 {
		return []*en.RolePermission{}, nil
	}

	input := make([]roleRefDTO, 0, len(roleRefs))

	for _, r := range roleRefs {
		input = append(input, roleRefDTO{
			RoleCode: r.RoleCode(),
			ScopeID:  r.ScopeID(),
		})
	}

	payload, err := json.Marshal(input)
	if err != nil {
		return nil, apperrors.New(core.JSON_SERIALIZATION_FAILED,
			apperrors.WithCauseDetail(err))
	}

	rows, err := s.queries.GetRolePermissionsByRoleRefs(ctx, payload)
	if err != nil {
		return nil, dberr.TranslateDBError(err, s.dberr)
	}

	result := make([]*en.RolePermission, 0, len(rows))

	for _, row := range rows {

		roleRef := vo.RestoreRoleRef(
			row.RoleCode,
			pg.UUIDPtrFromPgUUID(row.RoleScopeID),
		)

		role := en.RestoreRole(en.NewRoleParams{
			ID:              row.RoleID,
			RoleRef:         roleRef,
			RoleScopeType:   enums.RoleScopeType(row.RoleScopeType),
			Name:            row.RoleName,
			RoleAccessScope: enums.RoleAccessScope(row.RoleAccessScope),
			Level:           pg.Uint8FromInt32(row.RoleLevel),
			Description:     nil,
			IsSystem:        row.RoleIsSystem,
			IsSuper:         row.RoleIsSuper,
			IsActive:        row.RoleIsActive,
		})

		perm := en.RestorePermission(en.NewPermissionParams{
			ID:          row.RoleID,
			Code:        row.PermissionCode,
			Name:        pg.StringPtrFromText(row.PermissionName),
			Description: pg.StringPtrFromText(row.PermissionDescription),
			Resource:    row.PermissionResource,
			Action:      row.PermissionAction,
			IsActive:    row.PermissionIsActive,
		})

		result = append(result, en.NewRolePermission(role, perm))
	}

	return result, nil
}

// Update implements [repositories.RolePermissionRepository].
func (s *SqlcRolePermissionRepository) Update(ctx context.Context, rolePermission *en.RolePermission) (*en.RolePermission, *apperrors.AppError) {
	panic("unimplemented")
}
