-- name: GetRoleScopesByUserID :many
SELECT
    role_id,
    scope_id,
    scope_type
FROM role_assignment
WHERE user_id = $1;