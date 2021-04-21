BEGIN;

ALTER TABLE users ADD created_at DATETIME NULL;
ALTER TABLE users ADD updated_at DATETIME NULL;
ALTER TABLE roles ADD created_at DATETIME NULL;
ALTER TABLE roles ADD updated_at DATETIME NULL;
ALTER TABLE permissions ADD created_at DATETIME NULL;
ALTER TABLE permissions ADD updated_at DATETIME NULL;
ALTER TABLE role_permission ADD created_at DATETIME NULL;
ALTER TABLE role_permission ADD updated_at DATETIME NULL;

COMMIT;