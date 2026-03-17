-- +goose Up
SELECT 'up SQL query';

CREATE TABLE message (
    id TEXT NOT NULL PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT,
    deleted_at TEXT,
    author_id TEXT NOT NULL,
    is_readed INTEGER DEFAULT 0,
    content TEXT NOT NULL,
    content_type TEXT NOT NULL DEFAULT 'text' CHECK(content_type IN ('text', 'image', 'video'))
);

CREATE INDEX idx_message_content ON message(content);
CREATE INDEX idx_message_author_id ON message(author_id);
CREATE INDEX idx_message_created_at ON message(created_at);


-- +goose Down
SELECT 'down SQL query';

DROP TABLE message;