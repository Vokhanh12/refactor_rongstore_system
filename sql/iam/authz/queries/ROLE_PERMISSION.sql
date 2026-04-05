-- name: GetRolePermissionsByRoleRefs :many
SELECT 
    -- ROLE
    r.id            AS role_id,
    r.code          AS role_code,
    r.scope_id      AS role_scope_id,
    r.scope_type    AS role_scope_type,
    r.name          AS role_name,
    r.description   AS role_description,
    r.access_scope  AS role_access_scope,
    r.level         AS role_level,
    r.is_system     AS role_is_system,
    r.is_active     AS role_is_active,

    -- PERMISSION
    p.id            AS permission_id,       
    p.code          AS permission_code,
    p.name          AS permission_name,
    p.description   AS permission_description,
    p.resource      AS permission_resource,
    p.action        AS permission_action,
    p.is_active     AS permission_is_active

FROM roles r
JOIN role_permissions rp ON rp.role_id = r.id
JOIN permissions p ON p.id = rp.permission_id

JOIN jsonb_to_recordset($1::jsonb)
    AS x(role_code text, scope_id uuid)
ON r.code = x.role_code
AND r.scope_id = x.scope_id;