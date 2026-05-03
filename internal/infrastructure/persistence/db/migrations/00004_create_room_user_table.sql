-- +goose Up
SELECT 'up SQL query';

CREATE TABLE room_user(
    room_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    PRIMARY KEY (room_id, user_id),
    FOREIGN KEY (room_id) REFERENCES room(id) ON DELETE RESTRICT,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT
);

-- +goose Down
SELECT 'down SQL query';

DROP TABLE room_user;