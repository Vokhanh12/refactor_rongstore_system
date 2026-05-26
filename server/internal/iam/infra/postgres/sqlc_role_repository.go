package postgres

import (
	"context"

	"github.com/google/uuid"

	en "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	enu "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/postgres/mapper"

	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	sqlc "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ repositories.RoleRepository = (*SqlcRoleRepository)(nil)

type SqlcRoleRepository struct {
	dba *pg.DbAdapter
}

func NewSqlcRoleRepository(dba *pg.DbAdapter) repositories.RoleRepository {
	return &SqlcRoleRepository{
		dba: dba,
	}
}

// FindById implements [repositories.RoleRepository].
func (r *SqlcRoleRepository) FindById(ctx context.Context, id uuid.UUID) (*en.Role, *aerrs.AppError) {
	panic("unimplemented")
}

func (r *SqlcRoleRepository) Create(
	ctx context.Context,
	role *en.Role,
) (*en.Role, *aerrs.AppError) {

	row, err := r.dba.Q.CreateRole(
		ctx,
		sqlc.CreateRoleParams{
			ID:              role.ID(),
			ScopeID:         role.RoleKey().ScopeID(),
			RoleScopeType:   mapper.RoleScopeTypeToDB(role.ScopeType()),
			Code:            role.RoleKey().RoleCode(),
			Name:            role.Name(),
			Description:     role.Description(),
			RoleAccessScope: mapper.RoleAccessScopeToDB(role.AccessScope()),
			Level:           int32(role.Level()),
			IsSystem:        role.IsSystem(),
			IsActive:        role.IsActive(),
			IsSuper:         role.IsSuper(),
		},
	)

	if err != nil {
		return nil, r.dba.Wrap(err)
	}

	entity := mapper.CreateRoleRowToEntity(row)

	return &entity, nil
}

func (r *SqlcRoleRepository) Update(
	ctx context.Context,
	role *en.Role,
) (*en.Role, *aerrs.AppError) {

	row, err := r.dba.Q.UpdateRole(
		ctx,
		sqlc.UpdateRoleParams{
			ID:              role.ID(),
			ScopeID:         role.RoleKey().ScopeID(),
			RoleScopeType:   mapper.RoleScopeTypeToDB(role.ScopeType()),
			Code:            role.RoleKey().RoleCode(),
			Name:            role.Name(),
			Description:     role.Description(),
			RoleAccessScope: mapper.RoleAccessScopeToDB(role.AccessScope()),
			Level:           int32(role.Level()),
			IsSystem:        role.IsSystem(),
			IsActive:        role.IsActive(),
			IsSuper:         role.IsSuper(),
		},
	)

	if err != nil {
		return nil, r.dba.Wrap(err)
	}

	entity := mapper.UpdateRoleRowToEntity(row)

	return &entity, nil
}

func (r *SqlcRoleRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) *aerrs.AppError {

	err := r.dba.Q.DeleteRole(ctx, id)

	if err != nil {
		return r.dba.Wrap(err)
	}

	return nil
}

func (r *SqlcRoleRepository) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*en.Role, *aerrs.AppError) {
	panic("unimplemented")
}

func (r *SqlcRoleRepository) FindByCode(
	ctx context.Context,
	code string,
) (*en.Role, *aerrs.AppError) {
	panic("unimplemented")
}

func (r *SqlcRoleRepository) Exists(
	ctx context.Context,
	scopeType enu.RoleScopeType,
	roleKey vo.RoleKey,
) (bool, *aerrs.AppError) {

	exists, err := r.dba.Q.ExistsRoleByCodeScope(
		ctx,
		sqlc.ExistsRoleByCodeScopeParams{
			Code:          roleKey.RoleCode(),
			RoleScopeType: mapper.RoleScopeTypeToDB(scopeType),
			ScopeID:       roleKey.ScopeID(),
		},
	)

	if err != nil {
		return false, r.dba.Wrap(err)
	}

	return exists, nil
}
