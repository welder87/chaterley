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
// Возвращает ошибку core.ValidationError, если какое-то из полей невалидно.
func NewGroup(name string) (*Group, error) {
	newName, err := core.NewName(name)
	if err != nil {
		return nil, core.ValidationError{
			Field:  "Name",
			Reason: "Is empty",
			Err:    err,
		}
	}
	return &Group{
		id:        core.NewEntityID(),
		name:      newName,
		createdAt: core.NewCreatedAt(),
	}, nil
}

// ID - геттер для получения идентификатора группы пользователя чата.
func (g *Group) ID() core.EntityID {
	return g.id
}

// Name - геттер для получения наименования группы пользователя чата.
func (g *Group) Name() core.Name {
	return g.name
}

// SetName - сеттер для присваивания нового наименования группы пользователю чата.
func (g *Group) SetName(name string) error {
	newName, err := core.NewName(name)
	if err != nil {
		return core.ValidationError{
			Field:  "Name",
			Reason: "Is empty",
			Err:    err,
		}
	}
	if g.name == newName {
		return core.ValidationError{
			Field:  "Name",
			Reason: "Name unchanged",
			Err:    core.ErrNameUnchanged,
		}
	}
	g.name = newName
	g.updatedAt = core.NewUpdatedAt()
	return nil
}

// CreatedAt - геттер для получения даты создания группы пользователя чата.
func (g *Group) CreatedAt() core.CreatedAt {
	return g.createdAt
}

// UpdatedAt - геттер для получения даты обновления группы пользователя чата.
func (g *Group) UpdatedAt() core.UpdatedAt {
	return g.updatedAt
}

// DeletedAt - геттер для получения даты удаления группы пользователя чата.
func (g *Group) DeletedAt() core.DeletedAt {
	return g.deletedAt
}

func (g *Group) Delete() error {
	g.deletedAt = core.NewDeletedAt()
	g.updatedAt = core.NewUpdatedAt()
	return nil
}
