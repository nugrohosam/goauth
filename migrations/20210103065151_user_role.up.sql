BEGIN;

CREATE TABLE IF NOT EXISTS user_role (
    id serial PRIMARY KEY,
    user_id bigint,
    role_id bigint,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

COMMIt;