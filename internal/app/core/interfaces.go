package core

import (
	"context"
	"fmt"
)

type Valuer[Value comparable] interface {
	Val() Value
	Equal(other Value) bool
	fmt.Stringer
}

// Snapshooter - интерфейс для снимка состояния агрегата
type Snapshooter[Snapshot any] interface {
	// ToSnapshot возвращает snapshot
	ToSnapshot() Snapshot

	// FromSnapshot восстанавливает состояние из snapshot
	FromSnapshot(snapshot Snapshot) error
}

type Repository[Entity any] interface {
	Save(ctx context.Context, entity *Entity) error
	Remove(ctx context.Context, entity *Entity) error
	Get(ctx context.Context, entityID EntityID[Entity]) (*Entity, error)
	Exists(ctx context.Context, entityID EntityID[Entity]) (bool, error)
	ExistsIds(
		ctx context.Context,
		entityIDs []EntityID[Entity],
	) (map[EntityID[Entity]]struct{}, error)
}
