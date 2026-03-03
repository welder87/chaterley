package entities

import (
	"chaterley/internal/app/core"
)

type UserID core.EntityID[User]
type Login core.Login[User]
type Password core.PasswordHash[User]
type CreatedAt core.CreatedAt[User]
type UpdatedAt core.UpdatedAt[User]
type DeletedAt core.DeletedAt[User]

type User struct {
	id        UserID
	login     Login
	password  Password
	createdAt CreatedAt
	updatedAt UpdatedAt
	deletedAt DeletedAt
}
