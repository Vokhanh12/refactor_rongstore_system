package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/serialization"
	en "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/postgres/mapper"
	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ re.RoleRepository = (*SqlcRoleRepository)(nil)

type SqlcRoleRepository struct {
	queries *db.Queries
	dberr   dberr.DBError
}

func NewSqlcRoleRepository(queries *db.Queries, dberr dberr.DBError) repositories.RoleRepository {
	return &SqlcRoleRepository{queries: queries, dberr: dberr}
}

// Create implements [repositories.RoleRepository].
func (s *SqlcRoleRepository) Create(ctx context.Context, role *en.Role) (*en.Role, *apperrors.AppError) {

	createdRecord, err := s.queries.CreateRole(
		ctx, mapper.RoleToCreateParams(role),
	)

	if err != nil {
		return nil, dberr.TranslateDBError(err, s.dberr)
	}

	entity := mapper.CreateRoleRowToEntity(createdRecord)

	return &entity, nil
}

// Delete implements [repositories.RoleRepository].
func (s *SqlcRoleRepository) Delete(ctx context.Context, id uuid.UUID) *apperrors.AppError {
	err := s.queries.DeleteRole(ctx, id)

	if err != nil {
		return dberr.TranslateDBError(err, s.dberr)
	}

	return nil
}

// Update implements [repositories.RoleRepository].
func (s *SqlcRoleRepository) Update(ctx context.Context, role *en.Role) (*en.Role, *apperrors.AppError) {
	updatedRecord, err := s.queries.UpdateRole(
		ctx, mapper.RoleToUpdateParams(role),
	)

	if err != nil {
		return nil, dberr.TranslateDBError(err, s.dberr)
	}

	entity := mapper.UpdateRoleRowToEntity(updatedRecord)

	return &entity, nil
}

// FindByCode implements [repositories.RoleRepository].
func (s *SqlcRoleRepository) FindByCode(ctx context.Context, code string) (*en.Role, *apperrors.AppError) {
	panic("unimplemented")
}

// FindById implements [repositories.RoleRepository].
func (s *SqlcRoleRepository) FindById(ctx context.Context, id uuid.UUID) (*en.Role, *apperrors.AppError) {
	panic("unimplemented")
}

// Exists implements [repositories.RoleRepository].
func (s *SqlcRoleRepository) Exists(ctx context.Context, roleScopeType enums.RoleScopeType, roleRef valueobjects.RoleRef) (bool, *apperrors.AppError) {
	allowed, err := s.queries.ExistsRoleByCodeScope(ctx, mapper.RoleToExistsByCodeScopeParams(roleScopeType, roleRef))
	if err != nil {
		return false, dberr.TranslateDBError(err, s.dberr)
	}

	return allowed, nil
}

func (s *SqlcRoleRepository) ListRoleByRef(
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

	payload, err := serialization.MustMarshal(input)
	if err != nil {
		return nil, err
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

		role := en.RestoreRole(
			row.RoleID,
			en.RolePayload{
				RoleRef:         roleRef,
				RoleScopeType:   enums.RoleScopeType(row.RoleScopeType),
				Name:            row.RoleName,
				RoleAccessScope: enums.RoleAccessScope(row.RoleAccessScope),
				Level:           row.RoleLevel,
				Description:     pg.StringPtrFromText(row.RoleDescription),
				IsSystem:        row.RoleIsSystem,
				IsSuper:         row.RoleIsSuper,
				IsActive:        row.RoleIsActive,
			},
		)

		perm := en.RestorePermission(
			en.RestorePermissionParams{
				ID: row.RoleID,
				PermissionPayload: en.PermissionPayload{
					Code:        row.PermissionCode,
					Name:        pg.StringPtrFromText(row.PermissionName),
					Description: pg.StringPtrFromText(row.PermissionDescription),
					Resource:    row.PermissionResource,
					Action:      row.PermissionAction,
					IsActive:    row.PermissionIsActive,
				},
			},
		)

		result = append(result, en.NewRolePermission(role, perm))
	}

	return result, nil
}
