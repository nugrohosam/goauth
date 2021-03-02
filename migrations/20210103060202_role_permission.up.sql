BEGIN;

CREATE TABLE IF NOT EXISTS role_permission (
    id SERIAL,
    role_id BIGINT,
    permission_id BIGINT,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE,    
    PRIMARY KEY (id)

);

COMMIT;