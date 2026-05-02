package message

import (
	"chaterley/internal/app/room"
	"context"
)

type LastMessage struct {
	ID          string
	CreatedAt   string
	AuthorLogin string
	Content     string
}

type LoadLastMessagesByRoom func(
	ctx context.Context,
	roomID room.RoomID,
) ([]LastMessage, error)
