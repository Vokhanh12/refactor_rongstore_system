package postgres

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/testhelpers"
)

func TestSqlcRoleRepository_Create(t *testing.T) {
	testDB := testhelpers.SetupTestDB(t)
	defer testDB.Close(t)

	ctx := context.Background()

	repo := NewSqlcRoleRepository(testDB.Queries, dberr.RoleDbError)

	scopeID := uuid.New()

	roleRef := vo.NewRoleRefFromPersistence(
		"SUPER_ADMIN",
		&scopeID,
	)

	desc := "Super admin role"

	role := entities.NewRoleFromPersistence(
		uuid.New(),
		roleRef,
		entities.RoleTypeTenant,
		"Super Admin",
		entities.AcessScopeGobal,
		&desc,
		true,
		true,
		true,
	)

	created, err := repo.Create(ctx, &role)

	require.NoError(t, err)
	require.NotNil(t, created)

	assert.Equal(t, role.RoleRef().RoleCode(), created.RoleRef().RoleCode())
	assert.Equal(t, role.RoleRef().ScopeID(), created.RoleRef().ScopeID())
	assert.Equal(t, role.IsElevated(), created.IsElevated())
}

func TestSqlcRoleRepository_Update(t *testing.T) {

}

func TestSqlcRoleRepository_Delete(t *testing.T) {

}
