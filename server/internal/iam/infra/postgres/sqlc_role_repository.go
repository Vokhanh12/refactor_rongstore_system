package postgres

import (
	"context"

	"github.com/google/uuid"
	en "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/postgres/mapper"
	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
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
func (s *SqlcRoleRepository) Exists(ctx context.Context, roleScopeType enums.RoleScopeType, RoleKey valueobjects.RoleKey) (bool, *apperrors.AppError) {
	allowed, err := s.queries.ExistsRoleByCodeScope(ctx, mapper.RoleToExistsByCodeScopeParams(roleScopeType, RoleKey))
	if err != nil {
		return false, dberr.TranslateDBError(err, s.dberr)
	}

	return allowed, nil
}
