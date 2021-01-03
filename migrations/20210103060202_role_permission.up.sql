BEGIN;

CREATE TABLE IF NOT EXISTS role_permission (
    id bigint PRIMARY KEY AUTO_INCREMENT,
    role_id bigint,
    permission_id bigint,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);

COMMIt;