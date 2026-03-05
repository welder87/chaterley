package entities

import (
	"chaterley/internal/app/core"
)

type Group struct {
	id        core.EntityID
	name      core.Name
	createdAt core.CreatedAt
	updatedAt core.UpdatedAt
	deletedAt core.DeletedAt
}

func NewGroup(login string, password string) *Group {
	return &Group{
		id:        core.NewEntityID(),
		name:      core.NewName(login),
		createdAt: core.NewCreatedAt(),
	}
}
