package entities

import (
	"chaterley/internal/app/core"
)

const maxUserCount int = 100

// Room представляет комнату чата для изменения.
// Это агрегат.
type Room struct {
	// id - Идентификатор комнаты
	id core.EntityID
	// name - Наименование комнаты
	name core.Name
	// createdAt - Дата создания комнаты
	createdAt core.CreatedAt
	// updatedAt - Дата обновления комнаты
	updatedAt core.UpdatedAt
	// deletedAt - Дата удаления комнаты
	deletedAt core.DeletedAt
	// members - Пользователи в комнате.
	members map[core.EntityID]User
}

// NewRoom создает новую комнату.
// Возвращает ошибку core.ValidationError, если какое-то из полей невалидно.
func NewRoom(name string, adderID core.EntityID) (*Room, error) {
	newName, err := core.NewName(name)
	if err != nil {
		return nil, core.ValidationError{
			Field:  "Name",
			Reason: "Is empty",
			Err:    err,
		}
	}
	return &Room{
		id:        core.NewEntityID(),
		name:      newName,
		createdAt: core.NewCreatedAt(),
		members:   map[core.EntityID]User{},
	}, nil
}

// ID - геттер для получения идентификатора комнаты чата.
func (g *Room) ID() core.EntityID {
	return g.id
}

// Name - геттер для получения наименования комнаты чата.
func (g *Room) Name() core.Name {
	return g.name
}

// SetName - сеттер для присваивания нового наименования комнаты чата.
func (g *Room) SetName(name string, memberID core.EntityID) error {
	if _, ok := g.members[memberID]; !ok {
		return core.PermissionError{
			Field:  "Name",
			Reason: "Permission",
			Err:    core.ErrMemberNotFound,
		}
	}
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

// AddMember
func (g *Room) AddMember(adder, member User) error {
	if adder.id == member.id {
		return nil
	}
	if _, ok := g.members[adder.id]; !ok {
		return core.PermissionError{
			Field:  "Name",
			Reason: "Permission",
			Err:    core.ErrMemberNotFound,
		}
	}
	if len(g.members) > maxUserCount {
		return core.ValidationError{
			Field:  "Members",
			Reason: "Max User Count",
			Err:    core.ErrMemberCount,
		}
	}
	return nil
}

func (g *Room) RemoveMember(removerID core.EntityID, member User) error {
	if removerID == member.id {
		return core.PermissionError{
			Field:  "Members",
			Reason: "Permission",
			Err:    core.ErrMemberNotFound,
		}
	}
	if _, ok := g.members[removerID]; !ok {
		return core.PermissionError{
			Field:  "Members",
			Reason: "Permission",
			Err:    core.ErrMemberNotFound,
		}
	}
	return nil
}

// CreatedAt - геттер для получения даты создания комнаты пользователя чата.
func (g *Room) CreatedAt() core.CreatedAt {
	return g.createdAt
}

// UpdatedAt - геттер для получения даты обновления комнаты пользователя чата.
func (g *Room) UpdatedAt() core.UpdatedAt {
	return g.updatedAt
}

// DeletedAt - геттер для получения даты удаления комнаты пользователя чата.
func (g *Room) DeletedAt() core.DeletedAt {
	return g.deletedAt
}

func (g *Room) Delete() error {
	g.deletedAt = core.NewDeletedAt()
	g.updatedAt = core.NewUpdatedAt()
	return nil
}
