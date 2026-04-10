package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/adapter/mapper"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
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
func (s *SqlcRoleRepository) Create(ctx context.Context, role *entities.Role) (*entities.Role, *apperrors.AppError) {

	createdRecord, err := s.queries.CreateRole(
		ctx, mapper.CreateRoleToDBParams(role),
	)

	if err != nil {
		return nil, dberr.TranslateDBError(err, s.dberr)
	}

	entity := mapper.CreateRoleDBToRoleEntity(createdRecord)

	return &entity, nil

}

// Delete implements [repositories.RoleRepository].
func (s *SqlcRoleRepository) Delete(ctx context.Context, id uuid.UUID) *apperrors.AppError {
	panic("unimplemented")
}

// Update implements [repositories.RoleRepository].
func (s *SqlcRoleRepository) Update(ctx context.Context, role *entities.Role) (*entities.Role, *apperrors.AppError) {
	panic("unimplemented")
}
