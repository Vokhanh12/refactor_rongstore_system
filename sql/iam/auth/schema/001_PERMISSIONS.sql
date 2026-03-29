-- +goose Up
-- +goose StatementBegin

CREATE TABLE PERMISSIONS (
  ID UUID PRIMARY KEY,

  CODE VARCHAR(150) UNIQUE NOT NULL,

  NAME VARCHAR(255),
  DESCRIPTION TEXT,

  RESOURCE VARCHAR(100) NOT NULL,
  ACTION VARCHAR(50) NOT NULL,

  IS_ACTIVE BOOLEAN DEFAULT TRUE,

  CREATED_AT TIMESTAMP DEFAULT NOW(),
  UPDATED_AT TIMESTAMP DEFAULT NOW(),
  CREATED_BY UUID,
  UPDATED_BY UUID,

  CONSTRAINT UQ_PERMISSION_RESOURCE_ACTION
  UNIQUE (RESOURCE, ACTION)
);

CREATE INDEX IDX_PERMISSION_RESOURCE
ON PERMISSIONS(RESOURCE);

--   -- PERMISSIONS
-- INSERT INTO PERMISSIONS (ID, CODE, NAME, DESCRIPTION, RESOURCE, ACTION)
-- VALUES
--   (gen_random_uuid(), 'PERMISSION_CREATE', 'Create Permission', 'Create a new permission', 'PERMISSION', 'CREATE'),
--   (gen_random_uuid(), 'PERMISSION_UPDATE', 'Update Permission', 'Update permission info', 'PERMISSION', 'UPDATE'),
--   (gen_random_uuid(), 'PERMISSION_DELETE', 'Delete Permission', 'Delete a permission', 'PERMISSION', 'DELETE'),
--   (gen_random_uuid(), 'ORDER_CREATE', 'Create Order', 'User can create new orders', 'ORDER', 'CREATE'),
--   (gen_random_uuid(), 'ORDER_UPDATE', 'Update Order', 'User can update orders', 'ORDER', 'UPDATE'),
--   (gen_random_uuid(), 'ORDER_COMPLETE', 'Complete Order', 'User can mark order as complete', 'ORDER', 'COMPLETE'),
--   (gen_random_uuid(), 'PRODUCT_MANAGE', 'Manage Product', 'User can create/update products', 'PRODUCT', 'MANAGE'),
--   (gen_random_uuid(), 'PAYMENT_COLLECT', 'Collect Payment', 'User can collect payments', 'PAYMENT', 'COLLECT'),
--   (gen_random_uuid(), 'POST_CREATE', 'Create Post', 'User can create social posts', 'POST', 'CREATE');

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS PERMISSIONS;
-- +goose StatementEnd