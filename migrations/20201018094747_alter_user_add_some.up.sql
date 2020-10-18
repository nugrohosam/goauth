BEGIN;

ALTER TABLE users ADD COLUMN email varchar(255) NULL;
ALTER TABLE users ADD COLUMN password varchar(255) NULL;
ALTER TABLE users ADD COLUMN username varchar(255) NULL;
ALTER TABLE users ADD COLUMN phonenumber varchar(255) NULL;

COMMIt;