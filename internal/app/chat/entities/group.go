package entities

import (
	"chaterley/internal/app/core"
	"fmt"
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
// Возвращает ошибку, если какое-то из полей невалидно.
func NewGroup(name string) (*Group, error) {
	newName, err := core.NewName(name)
	if err != nil {
		return nil, fmt.Errorf("Error with group name: %w", err)
	}
	return &Group{
		id:        core.NewEntityID(),
		name:      newName,
		createdAt: core.NewCreatedAt(),
	}, nil
}
