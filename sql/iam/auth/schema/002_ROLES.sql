-- +goose Up
-- +goose StatementBegin

CREATE TABLE ROLES (
  ID UUID PRIMARY KEY,

  CODE VARCHAR(100) UNIQUE NOT NULL,
  NAME VARCHAR(255),

  TYPE VARCHAR(30) NOT NULL, -- PLATFORM | ORGANIZATION
  DESCRIPTION TEXT,

  IS_SYSTEM BOOLEAN DEFAULT FALSE,
  IS_SUPER BOOLEAN DEFAULT FALSE,
  IS_ACTIVE BOOLEAN DEFAULT TRUE,

  CREATED_AT TIMESTAMP DEFAULT NOW(),
  UPDATED_AT TIMESTAMP DEFAULT NOW(),
  CREATED_BY UUID,
  UPDATED_BY UUID
);

CREATE INDEX IDX_ROLES_TYPE
ON ROLES(TYPE);


-- -- ROLES
-- INSERT INTO ROLES (ID, CODE, NAME, TYPE, DESCRIPTION, IS_SYSTEM, IS_SUPER)
-- VALUES
--   (gen_random_uuid(), 'SUPER_ADMIN', 'Super Admin', 'PLATFORM', 'Full access to system', TRUE, TRUE),
--   (gen_random_uuid(), 'RESTAURANT_MANAGER', 'Restaurant Manager', 'ORGANIZATION', 'Manage a restaurant unit', FALSE, FALSE),
--   (gen_random_uuid(), 'CAFE_OPERATOR', 'Café Operator', 'ORGANIZATION', 'Operate a café', FALSE, FALSE),
--   (gen_random_uuid(), 'FB_VENDOR', 'Street Vendor', 'ORGANIZATION', 'Operate a mobile/outlet vendor', FALSE, FALSE);

ON CONFLICT (CODE) DO NOTHING;

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ROLES;
-- +goose StatementEnd