-- name: GetRoleScopesByUserID :many
SELECT
    role_id,
    scope_id,
    scope_type
FROM ROLE_ASSIGNMENTS
WHERE user_id = $1;