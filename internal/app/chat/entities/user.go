package entities

import (
	"chaterley/internal/app/core"
	"time"
)

// User представляет пользователя чата.
type User struct {
	// id - Идентификатор пользователя
	id core.EntityID
	// login - Логин пользователя
	login core.Login
	// createdAt - Дата создания пользователя
	createdAt core.CreatedAt
	// updatedAt - Дата обновления пользователя
	updatedAt core.UpdatedAt
	// deletedAt - Дата удаления пользователя
	deletedAt *core.DeletedAt
}

// NewUser создает нового пользователя.
// В дальнейшем должна возвращать ошибку, если какое-то из полей невалидно.
func NewUser(login string) *User {
	return &User{
		id:        core.NewEntityID(),
		login:     core.NewLogin(login),
		createdAt: core.NewCreatedAt(),
		updatedAt: core.NewUpdatedAt(),
	}
}

func (u *User) ToSnapshot() UserSnapshot {
	return UserSnapshot{
		ID:        u.id.Val(),
		Login:     u.login.Val(),
		CreatedAt: u.createdAt.Val(),
		UpdatedAt: u.updatedAt.Val(),
		DeletedAt: u.deletedAt.Val(),
	}
}

type UserSnapshot struct {
	// ID - Идентификатор пользователя
	ID string
	// Login - Логин пользователя
	Login string
	// CreatedAt - Дата создания пользователя
	CreatedAt time.Time
	// UpdatedAt - Дата обновления пользователя
	UpdatedAt time.Time
	// DeletedAt - Дата удаления пользователя
	DeletedAt time.Time
}
