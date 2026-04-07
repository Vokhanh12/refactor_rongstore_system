package postgres

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	iam_sqlc "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/testhelpers"
)

func TestSqlcRolePermissionRepository_FindAllByRoleRefs(t *testing.T) {
	ctx := context.Background()

	// setup DB (testcontainer hoặc test DB local)
	testDB := testhelpers.SetupTestDB(t)
	defer testDB.Close(t)

	queries := iam_sqlc.New(testDB.DB)
	repo := NewSqlcRolePermissionRepository(queries)

	// ===== Seed data =====
	roleID := uuid.New()
	permID := uuid.New()

	scopeID := uuid.New()

	// insert role
	_, err := testDB.DB.Exec(ctx, `
		INSERT INTO roles (
			id, code, role_scope_type, scope_id, name,
			role_access_scope, is_system, is_active
		) VALUES ($1,$2,'TENANT',$3,$4,'ALL',false,true)
	`, roleID, "ADMIN", scopeID, "Admin Role")
	require.NoError(t, err)

	// insert permission
	_, err = testDB.DB.Exec(ctx, `
		INSERT INTO permissions (
			id, code, name, resource, action, is_active
		) VALUES ($1,$2,$3,$4,$5,true)
	`, permID, "USER_CREATE", "Create User", "USER", "CREATE")
	require.NoError(t, err)

	// insert role_permission
	_, err = testDB.DB.Exec(ctx, `
		INSERT INTO role_permissions (role_id, permission_id)
		VALUES ($1,$2)
	`, roleID, permID)
	require.NoError(t, err)

	// ===== Call function =====
	roleRef := valueobjects.NewRoleRefFromPersistence(
		"ADMIN",
		&[]string{scopeID.String()}[0],
	)

	results, appErr := repo.FindAllByRoleRefs(ctx, []valueobjects.RoleRef{roleRef})

	// ===== Assertions =====
	require.Nil(t, appErr)
	require.Len(t, results, 1)

	rp := results[0]

	assert.Equal(t, "ADMIN", rp.Role.RoleRef().RoleCode())
	assert.Equal(t, scopeID.String(), rp.Role.RoleRef().ScopeID())

	assert.Equal(t, "USER_CREATE", rp.Permission.Code())
	assert.Equal(t, "USER", rp.Permission.ResourceAction().Resource())
	assert.Equal(t, "CREATE", rp.Permission.ResourceAction().Action())
}

func TestSqlcRolePermissionRepository_FindAllByRoleRefs_EmptyInput(t *testing.T) {
	ctx := context.Background()

	testDB := testhelpers.SetupTestDB(t)
	defer testDB.Close(t)

	queries := iam_sqlc.New(testDB.DB)
	repo := NewSqlcRolePermissionRepository(queries)

	results, err := repo.FindAllByRoleRefs(ctx, []valueobjects.RoleRef{})

	require.Nil(t, err)
	assert.Empty(t, results)
}

func TestSqlcRolePermissionRepository_FindAllByRoleRefs_NotFound(t *testing.T) {
	ctx := context.Background()

	testDB := testhelpers.SetupTestDB(t)
	defer testDB.Close(t)

	queries := iam_sqlc.New(testDB.DB)
	repo := NewSqlcRolePermissionRepository(queries)

	scopeID := uuid.New()

	roleRef := valueobjects.NewRoleRefFromPersistence(
		"NOT_EXIST",
		&[]string{scopeID.String()}[0],
	)

	results, err := repo.FindAllByRoleRefs(ctx, []valueobjects.RoleRef{roleRef})

	require.Nil(t, err)
	assert.Empty(t, results)
}
