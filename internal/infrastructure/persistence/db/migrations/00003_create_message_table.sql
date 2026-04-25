-- +goose Up
SELECT 'up SQL query';

CREATE TABLE message (
    id TEXT NOT NULL PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT DEFAULT NULL,
    deleted_at TEXT DEFAULT NULL,
    author_id TEXT NOT NULL,
    room_id TEXT NOT NULL,
    content TEXT NOT NULL,
    FOREIGN KEY(author_id) REFERENCES user(id) ON DELETE RESTRICT
    FOREIGN KEY(room_id) REFERENCES room(id) ON DELETE RESTRICT
);

CREATE INDEX idx_message_content ON message(content);
CREATE INDEX idx_message_author_id ON message(author_id);
CREATE INDEX idx_message_room_id ON message(room_id);

-- +goose Down
SELECT 'down SQL query';

DROP TABLE message;