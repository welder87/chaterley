package user

import (
	"chaterley/internal/app/core"
	"time"

	"github.com/google/uuid"
)

type (
	UserID       = core.EntityID[User]
	CreatedAt    = core.CreatedAt[User]
	UpdatedAt    = core.UpdatedAt[User]
	DeletedAt    = core.DeletedAt[User]
	Login        = core.Login[User]
	PasswordHash = core.PasswordHash[User]
)

// User представляет пользователя чата.
// Основная сущность доменной области Chat.
type User struct {
	// id - Идентификатор пользователя
	id UserID
	// login - Логин пользователя
	login Login
	// password - Хеш пароля пользователя
	password PasswordHash
	// createdAt - Дата создания пользователя
	createdAt CreatedAt
	// updatedAt - Дата обновления пользователя
	updatedAt UpdatedAt
	// deletedAt - Дата удаления пользователя
	deletedAt *DeletedAt
}

// NewUser создает нового пользователя.
// В дальнейшем должна возвращать ошибку, если какое-то из полей невалидно.
func NewUser(login string, password string) *User {
	return &User{
		id:        core.NewEntityID[User](),
		login:     core.NewLogin[User](login),
		password:  core.NewPasswordHash[User](password),
		createdAt: core.NewCreatedAt[User](),
		updatedAt: core.NewUpdatedAt[User](),
	}
}

func (u *User) ToSnapshot() UserSnapshot {
	snapshot := UserSnapshot{
		ID:        u.id.Val(),
		Login:     u.login.Val(),
		Password:  u.password.Val(),
		CreatedAt: u.createdAt.Val(),
		UpdatedAt: u.updatedAt.Val(),
	}
	if u.deletedAt != nil {
		deletedAt := u.deletedAt.Val()
		snapshot.DeletedAt = &deletedAt
	}
	return snapshot
}

type UserSnapshot struct {
	// ID - Идентификатор пользователя
	ID uuid.UUID
	// Login - Логин пользователя
	Login string
	// Password - Хеш пароля пользователя
	Password string
	// CreatedAt - Дата создания пользователя
	CreatedAt time.Time
	// UpdatedAt - Дата обновления пользователя
	UpdatedAt time.Time
	// DeletedAt - Дата удаления пользователя
	DeletedAt *time.Time
}
