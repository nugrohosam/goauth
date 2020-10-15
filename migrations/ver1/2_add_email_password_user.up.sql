BEGIN;

ALTER TABLE users 
ADD COLUMN email varchar(255) NOT NULL;

ALTER TABLE users 
ADD COLUMN password varchar(255) NOT NULL;

COMMIT;