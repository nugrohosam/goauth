BEGIN;

CREATE TABLE IF NOT EXISTS role_permission (
    id serial PRIMARY KEY,
    role_id bigint,
    permission_id bigint,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);

COMMIt;