-- ============================================================
-- WRITE: ROLE COMMANDS
-- ============================================================

-- name: CreateRole :one
INSERT INTO roles (
    id,
    scope_id,
    role_scope_type,
    code,
    name,
    description,
    role_access_scope,
    level,
    is_system,
    is_active,
    is_super
) VALUES (
    $1, $2, $3, $4, $5, $6,
    $7, $8, $9, $10, $11
)
RETURNING
    id,
    scope_id,
    role_scope_type,
    code,
    name,
    description,
    role_access_scope,
    level,
    is_system,
    is_active,
    is_super
;

-- name: UpdateRole :one
UPDATE roles
SET
    scope_id = $2,
    role_scope_type = $3,
    code = $4,
    name = $5,
    description = $6,
    role_access_scope = $7,
    level = $8,
    is_system = $9,
    is_active = $10,
    is_super = $11
WHERE id = $1
RETURNING
    id,
    scope_id,
    role_scope_type,
    code,
    name,
    description,
    role_access_scope,
    level,
    is_system,
    is_active,
    is_super
;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1;


-- ============================================================
-- READ: ROLE QUERIES
-- ============================================================

-- name: GetRoleByID :one
SELECT
    id,
    scope_id,
    role_scope_type,
    code,
    name,
    description,
    role_access_scope,
    level,
    is_system,
    is_active,
    is_super
FROM roles
WHERE id = $1;

-- name: GetRoleByCode :one
SELECT
    id,
    scope_id,
    role_scope_type,
    code,
    name,
    description,
    role_access_scope,
    level,
    is_system,
    is_active,
    is_super
FROM roles
WHERE code = $1;

-- name: ListRoles :many
SELECT
    id,
    scope_id,
    role_scope_type,
    code,
    name,
    description,
    role_access_scope,
    level,
    is_system,
    is_active,
    is_super
FROM roles
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ExistsRoleByCodeScope :one
SELECT EXISTS (
    SELECT 1
    FROM roles
    WHERE code = $1
      AND role_scope_type = $2
      AND scope_id IS NOT DISTINCT FROM $3
) AS exists;