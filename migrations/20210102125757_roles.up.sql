BEGIN;

CREATE TABLE IF NOT EXISTS roles (
    id serial PRIMARY KEY,
    name varchar(255) NULL
);

COMMIt;