BEGIN;

ALTER TABLE users DROP COLUMN created_at;
ALTER TABLE users DROP COLUMN updated_at;
ALTER TABLE roles DROP COLUMN created_at;
ALTER TABLE roles DROP COLUMN updated_at;
ALTER TABLE permissions DROP COLUMN created_at;
ALTER TABLE permissions DROP COLUMN updated_at;
ALTER TABLE role_permission DROP COLUMN created_at;
ALTER TABLE role_permission DROP COLUMN updated_at;

COMMIT;