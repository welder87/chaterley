package entities

import (
	"chaterley/internal/app/core"
)

// User представляет пользователя чата.
// Основная сущность доменной области Chat.
type User struct {
	// id - Идентификатор пользователя
	id core.EntityID
	// groupID - Группа пользователя
	groupID core.EntityID
	// login - Логин пользователя
	login core.Login
	// password - Хеш пароля пользователя
	password core.PasswordHash
	// createdAt - Дата создания пользователя
	createdAt core.CreatedAt
	// updatedAt - Дата обновления пользователя
	updatedAt core.UpdatedAt
	// deletedAt - Дата удаления пользователя
	deletedAt core.DeletedAt
}

// NewUser создает нового пользователя.
// В дальнейшем должна возвращать ошибку, если какое-то из полей невалидно.
func NewUser(login string, password string) *User {
	return &User{
		id:        core.NewEntityID(),
		login:     core.NewLogin(login),
		password:  core.NewPasswordHash(password),
		createdAt: core.NewCreatedAt(),
	}
}
