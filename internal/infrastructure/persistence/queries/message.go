package queries

import (
	"chaterley/internal/app/message"
	"chaterley/internal/app/room"
	"context"
	"database/sql"
	"fmt"
)

func NewLoadLastMessagesByRoom(
	readDbConn *sql.DB,
	lastMessageCount int,
) message.LoadLastMessagesByRoom {
	return func(ctx context.Context, roomID room.RoomID) ([]message.LastMessage, error) {
		query := `
		SELECT
			message.id,
			message.created_at,
			user.login,
			message.content
		FROM message
		JOIN user ON user.id = message.author_id
		WHERE room_id = ?
		LIMIT ?
		`
		stmt, err := readDbConn.PrepareContext(ctx, query)
		if err != nil {
			return []message.LastMessage{}, err
		}
		defer stmt.Close()
		rows, err := stmt.QueryContext(ctx, roomID.String(), lastMessageCount)
		if err != nil {
			return []message.LastMessage{}, fmt.Errorf("query error: %w", err)
		}
		defer rows.Close()
		messages := []message.LastMessage{}
		for rows.Next() {
			var msg message.LastMessage
			err = rows.Scan(
				&msg.ID,
				&msg.CreatedAt,
				&msg.AuthorLogin,
				&msg.Content,
			)
			if err != nil {
				return []message.LastMessage{}, err
			}
			messages = append(messages, msg)
		}
		return messages, nil
	}
}
