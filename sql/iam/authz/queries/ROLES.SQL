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