BEGIN;

CREATE TABLE IF NOT EXISTS users (
    id BIGINT AUTO_INCREMENT,
    name VARCHAR(255) DEFAULT NULL,
    PRIMARY KEY (id)
);

COMMIT;