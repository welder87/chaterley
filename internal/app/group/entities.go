package group

import (
	"chaterley/internal/app/core"
)

type (
	GroupID        = core.EntityID[Group]
	GroupCreatedAt = core.CreatedAt[Group]
	GroupUpdatedAt = core.UpdatedAt[Group]
	GroupDeletedAt = core.DeletedAt[Group]
	GroupName      = core.Name[Group]
)

// Group представляет группу пользователя чата.
// Отвечает за права доступа. Например, admin.
type Group struct {
	// id - Идентификатор группы
	id GroupID
	// name - Наименование группы
	name GroupName
	// createdAt - Дата создания группы
	createdAt GroupCreatedAt
	// updatedAt - Дата обновления группы
	updatedAt GroupUpdatedAt
	// deletedAt - Дата удаления группы
	deletedAt *GroupDeletedAt
}

// NewGroup создает новую группу.
// Возвращает ошибку core.ValidationError, если какое-то из полей невалидно.
func NewGroup(name string) (*Group, error) {
	newName, err := core.NewName[Group](name)
	if err != nil {
		return nil, core.ValidationError{Field: "name", Code: core.Unknown, Err: err}
	}
	return &Group{
		id:        core.NewEntityID[Group](),
		name:      newName,
		createdAt: core.NewCreatedAt[Group](),
		updatedAt: core.NewUpdatedAt[Group](),
	}, nil
}

// SetName - сеттер для присваивания нового наименования Группы Пользователя Чата.
func (g *Group) SetName(name string) error {
	newName, err := core.NewName[Group](name)
	if err != nil {
		return core.ValidationError{Field: "name", Code: core.Unknown, Err: err}
	}
	if g.name == newName {
		return core.ValidationError{
			Field: "name",
			Code:  core.NameUnchanged,
		}
	}
	g.name = newName
	g.updatedAt = core.NewUpdatedAt[Group]()
	return nil
}

func (g *Group) Delete() error {
	deletedAt := core.NewDeletedAt[Group]()
	g.deletedAt = &deletedAt
	g.updatedAt = core.NewUpdatedAt[Group]()
	return nil
}
