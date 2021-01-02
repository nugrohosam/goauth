BEGIN;

CREATE TABLE IF NOT EXISTS permissions (
    id bigint PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NULL
);

COMMIt;