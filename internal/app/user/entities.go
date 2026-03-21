package user

import (
	"chaterley/internal/app/core"
	"time"

	"github.com/google/uuid"
)

// User представляет пользователя чата.
// Основная сущность доменной области Chat.
type User struct {
	// id - Идентификатор пользователя
	id core.EntityID[User]
	// login - Логин пользователя
	login core.Login
	// password - Хеш пароля пользователя
	password core.PasswordHash
	// createdAt - Дата создания пользователя
	createdAt core.CreatedAt
	// updatedAt - Дата обновления пользователя
	updatedAt core.UpdatedAt
	// deletedAt - Дата удаления пользователя
	deletedAt *core.DeletedAt
}

// NewUser создает нового пользователя.
// В дальнейшем должна возвращать ошибку, если какое-то из полей невалидно.
func NewUser(login string, password string) *User {
	return &User{
		id:        core.NewEntityID[User](),
		login:     core.NewLogin(login),
		password:  core.NewPasswordHash(password),
		createdAt: core.NewCreatedAt(),
		updatedAt: core.NewUpdatedAt(),
	}
}

func (u *User) ToSnapshot() UserSnapshot {
	return UserSnapshot{
		ID:        u.id.Val(),
		Login:     u.login.Val(),
		Password:  u.password.Val(),
		CreatedAt: u.createdAt.Val(),
		UpdatedAt: u.updatedAt.Val(),
		DeletedAt: u.deletedAt.Val(),
	}
}

type UserSnapshot struct {
	// ID - Идентификатор пользователя
	ID uuid.UUID
	// Login - Логин пользователя
	Login string
	// Password - Хеш пароля пользователя
	Password core.PasswordHash
	// CreatedAt - Дата создания пользователя
	CreatedAt time.Time
	// UpdatedAt - Дата обновления пользователя
	UpdatedAt time.Time
	// DeletedAt - Дата удаления пользователя
	DeletedAt time.Time
}


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

func (g *Group) Delete() error {
	g.deletedAt = core.NewDeletedAt()
	g.updatedAt = core.NewUpdatedAt()
	return nil
}
