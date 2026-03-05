package entities

import (
	"chaterley/internal/app/core"
)

type User struct {
	id        core.EntityID
	groupID   core.EntityID
	login     core.Login
	password  core.PasswordHash
	createdAt core.CreatedAt
	updatedAt core.UpdatedAt
	deletedAt core.DeletedAt
}

func NewUser(login string, password string) *User {
	return &User{
		id:        core.NewEntityID(),
		login:     core.NewLogin(login),
		password:  core.NewPasswordHash(password),
		createdAt: core.NewCreatedAt(),
	}
}
