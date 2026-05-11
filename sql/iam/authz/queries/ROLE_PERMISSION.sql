-- name: ListRolePermissionByRoleRefs :many
SELECT
    r.id AS role_id,
    p.id AS permission_id,
    
FROM roles r

JOIN role_permissions rp
    ON rp.role_id = r.id

JOIN permissions p
    ON p.id = rp.permission_id

JOIN jsonb_to_recordset($1::jsonb)
    AS x(role_code text, scope_id uuid)

ON r.code = x.role_code

AND (
    r.scope_id = x.scope_id
    OR (r.scope_id IS NULL AND x.scope_id IS NULL)
);