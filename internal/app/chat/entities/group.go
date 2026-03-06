package entities

import (
	"chaterley/internal/app/core"
)

// Group представляет группу пользователя чата.
// Отвечает за права доступа. Например, admin.
type Group struct {
	// id - Идентификатор группы
	id core.EntityID
	// name - Наименование группы
	name core.Name
	// createdAt - Дата создания группы
	createdAt core.CreatedAt
	// updatedAt - Дата обновления группы
	updatedAt core.UpdatedAt
	// deletedAt - Дата удаления группы
	deletedAt core.DeletedAt
}

// NewGroup создает новую группу.
// В дальнейшем должна возвращать ошибку, если какое-то из полей невалидно.
func NewGroup(name string) *Group {
	return &Group{
		id:        core.NewEntityID(),
		name:      core.NewName(name),
		createdAt: core.NewCreatedAt(),
	}
}
