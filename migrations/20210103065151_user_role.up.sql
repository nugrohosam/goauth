BEGIN;

CREATE TABLE IF NOT EXISTS user_role (
    id bigint PRIMARY KEY AUTO_INCREMENT,
    user_id bigint,
    role_id bigint,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

COMMIt;