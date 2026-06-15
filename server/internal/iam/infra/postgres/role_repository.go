package postgres

import (
	"context"

	"github.com/google/uuid"

	en "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/postgres/mapper"

	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	sqlc "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ repositories.RoleRepository = (*RoleRepository)(nil)

type RoleRepository struct {
	dba *pg.DbAdapter
}

func NewRoleRepository(dba *pg.DbAdapter) repositories.RoleRepository {
	return &RoleRepository{
		dba: dba,
	}
}

func (r *RoleRepository) Create(
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
		return nil, r.dba.Translate(err)
	}

	entity := mapper.CreateRoleRowToEntity(row)

	return &entity, nil
}

func (r *RoleRepository) Update(
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
		return nil, r.dba.Translate(err)
	}

	entity := mapper.UpdateRoleRowToEntity(row)

	return &entity, nil
}

func (r *RoleRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) *aerrs.AppError {

	err := r.dba.Q.DeleteRole(ctx, id)

	if err != nil {
		return r.dba.Translate(err)
	}

	return nil
}
