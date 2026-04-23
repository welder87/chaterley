-- +goose Up
SELECT 'up SQL query';

CREATE TABLE room(
    id TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT DEFAULT NULL,
    deleted_at TEXT DEFAULT NULL
);

-- +goose Down
SELECT 'down SQL query';
