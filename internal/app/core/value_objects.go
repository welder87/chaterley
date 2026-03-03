package core

import (
	"time"

	"github.com/google/uuid"
)

type EntityID[T any] struct {
	val uuid.UUID
}

type Login[T any] struct {
	val string
}

type PasswordHash[T any] struct {
	val string
}

type CreatedAt[T any] struct {
	val time.Time
}

type UpdatedAt[T any] struct {
	val time.Time
}

type DeletedAt[T any] struct {
	val time.Time
}
