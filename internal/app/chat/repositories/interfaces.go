package repositories

import (
	"chaterley/internal/app/chat/entities"
	"chaterley/internal/app/core"
	"context"
)

type RoomRepository interface {
	Save(ctx context.Context, room *entities.Room) error
	Add(ctx context.Context, room *entities.Room) error
	Remove(ctx context.Context, room *entities.Room) error
	GetByID(ctx context.Context, roomID core.EntityID) (*entities.Room, error)
	GetMessages(
		ctx context.Context,
		roomID core.EntityID,
		beforeID core.EntityID,
		limit int,
	) ([]*entities.Message, core.EntityID, bool, error)
}
