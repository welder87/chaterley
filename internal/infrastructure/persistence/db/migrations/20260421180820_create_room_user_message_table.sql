-- +goose Up
SELECT 'up SQL query';

CREATE TABLE room_user_message(
    room_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    message_id TEXT NOT NULL,
    PRIMARY KEY (room_id, user_id, message_id),
    FOREIGN KEY (room_id) REFERENCES room(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
    FOREIGN KEY (message_id) REFERENCES message(id) ON DELETE CASCADE
);

-- +goose Down
SELECT 'down SQL query';
