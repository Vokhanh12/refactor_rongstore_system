package postgres

import (
	"context"

	"github.com/google/uuid"

	en "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	enu "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/postgres/mapper"

	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ repositories.RoleRepository = (*SqlcRoleRepository)(nil)

type SqlcRoleRepository struct {
	queries *db.Queries
	dberr   dberr.DBError
}

func NewSqlcRoleRepository(
	queries *db.Queries,
	dberr dberr.DBError,
) repositories.RoleRepository {
	return &SqlcRoleRepository{
		queries: queries,
		dberr:   dberr,
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

	row, err := r.queries.CreateRole(
		ctx,
		db.CreateRoleParams{
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
		return nil, dberr.TranslateDBError(err, r.dberr)
	}

	entity := mapper.CreateRoleRowToEntity(row)

	return &entity, nil
}

func (r *SqlcRoleRepository) Update(
	ctx context.Context,
	role *en.Role,
) (*en.Role, *aerrs.AppError) {

	row, err := r.queries.UpdateRole(
		ctx,
		db.UpdateRoleParams{
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
		return nil, dberr.TranslateDBError(err, r.dberr)
	}

	entity := mapper.UpdateRoleRowToEntity(row)

	return &entity, nil
}

func (r *SqlcRoleRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) *aerrs.AppError {

	err := r.queries.DeleteRole(ctx, id)

	if err != nil {
		return dberr.TranslateDBError(err, r.dberr)
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

	exists, err := r.queries.ExistsRoleByCodeScope(
		ctx,
		db.ExistsRoleByCodeScopeParams{
			Code:          roleKey.RoleCode(),
			RoleScopeType: mapper.RoleScopeTypeToDB(scopeType),
			ScopeID:       roleKey.ScopeID(),
		},
	)

	if err != nil {
		return false, dberr.TranslateDBError(err, r.dberr)
	}

	return exists, nil
}
