package postgres

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
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
func (s *SqlcRolePermissionRepository) Create(ctx context.Context, rolePermission *entities.RolePermission) (*entities.RolePermission, *apperrors.AppError) {
	panic("unimplemented")
}

// Delete implements [repositories.RolePermissionRepository].
func (s *SqlcRolePermissionRepository) Delete(ctx context.Context, id uuid.UUID) *apperrors.AppError {
	panic("unimplemented")
}

func (s *SqlcRolePermissionRepository) FindAllByRoleRefs(
	ctx context.Context,
	roleRefs []valueobjects.RoleRef,
) ([]*entities.RolePermission, *apperrors.AppError) {

	if len(roleRefs) == 0 {
		return []*entities.RolePermission{}, nil
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
		return nil, apperrors.New(domain.JSON_SERIALIZATION_FAILED,
			apperrors.WithCauseDetail(err))
	}

	rows, err := s.queries.GetRolePermissionsByRoleRefs(ctx, payload)
	if err != nil {
		return nil, dberr.TranslateDBError(err, s.dberr)
	}

	result := make([]*entities.RolePermission, 0, len(rows))

	for _, row := range rows {

		roleRef := valueobjects.NewRoleRefFromPersistence(
			row.RoleCode,
			pg.UUIDToString(row.RoleScopeID),
		)

		role := entities.NewRoleFromPersistence(
			row.RoleID.String(),
			roleRef,
			entities.RoleScopeType(row.RoleScopeType),
			row.RoleName.String,
			entities.RoleType(row.RoleType),
			row.RoleDescription.String,
			row.RoleIsSystem.Bool,
			row.RoleIsSuper.Bool,
			row.RoleIsActive,
		)

		perm := entities.NewPermissionFromPersistence(
			row.PermissionID.String(),
			row.PermissionCode,
			row.PermissionName,
			row.PermissionDescription,
			row.PermissionResource,
			row.PermissionAction,
			row.PermissionIsActive,
		)

		result = append(result, entities.NewRolePermission(role, perm))
	}
}

// Update implements [repositories.RolePermissionRepository].
func (s *SqlcRolePermissionRepository) Update(ctx context.Context, rolePermission *entities.RolePermission) (*entities.RolePermission, *apperrors.AppError) {
	panic("unimplemented")
}
