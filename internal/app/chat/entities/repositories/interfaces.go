package repositories

import (
	"chaterley/internal/app/core"
	"context"
)

type Repository[Entity any] interface {
	Save(ctx context.Context, entity *Entity) error
	Remove(ctx context.Context, entity *Entity) error
	Get(ctx context.Context, entityID core.EntityID) (*Entity, error)
}
