BEGIN;
ALTER TABLE users ADD COLUMN email VARCHAR(255) NULL;
ALTER TABLE users ADD COLUMN password VARCHAR(255) NULL;
ALTER TABLE users ADD COLUMN username VARCHAR(255) NULL;
ALTER TABLE users ADD COLUMN phonenumber VARCHAR(255) NULL;
COMMIT;
