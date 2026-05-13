-- name: ListAuthorizationGrantsByRoleKeys :many
SELECT
    r.code AS role_code,
    r.scope_id AS role_scope_id,
    r.is_super AS role_is_super,

    p.resource AS permission_resource,
    p.action AS permission_action

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