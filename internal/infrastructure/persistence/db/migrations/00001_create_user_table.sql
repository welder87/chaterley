-- +goose Up
SELECT 'up SQL query';

CREATE TABLE user(
    id TEXT NOT NULL PRIMARY KEY,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    password_salt TEXT NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT DEFAULT NULL,
    deleted_at TEXT DEFAULT NULL
);

CREATE INDEX idx_user_login ON user(login);
CREATE INDEX idx_user_password ON user(password);

-- +goose Down
SELECT 'down SQL query';

DROP TABLE user;
