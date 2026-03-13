package repositories

// Snapshooter - типобезопасный интерфейс с дженериком
type Snapshooter[Snapshot any] interface {
	// ToSnapshot возвращает snapshot
	ToSnapshot() Snapshot

	// FromSnapshot восстанавливает состояние из snapshot
	FromSnapshot(snapshot Snapshot) error
}

type Repository[Snapshot any] interface {
	save(data Snapshooter[Snapshot])
}
