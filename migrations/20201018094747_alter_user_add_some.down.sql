BEGIN;

ALTER TABLE users 
    DROP COLUMN email,
    DROP COLUMN password,
    DROP COLUMN username,
    DROP COLUMN phonenumber;

COMMIT;