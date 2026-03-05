package core

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type EntityID struct {
	val uuid.UUID
}

func NewEntityID() EntityID {
	id, err := uuid.NewV7()
	if err != nil {
		// считаем, что ситуация совсем неадекватная, и падаем
		panic(err)
	}
	return EntityID{
		val: id,
	}
}

type Login struct {
	val string
}

func NewLogin(val string) Login {
	val = strings.TrimSpace(val)
	// тут должны быть проверки (бизнес-правила для логина)
	return Login{val: val}
}

type Name struct {
	val string
}

func NewName(val string) Name {
	val = strings.TrimSpace(val)
	// тут должны быть проверки (бизнес-правила для имени)
	return Name{val: val}
}

type PasswordHash struct {
	val string
}

func NewPasswordHash(password string) PasswordHash {
	val := strings.TrimSpace(password)
	// тут должны быть проверки (бизнес-правила для пароля) с хешированием
	// мы не должны хранить входной val, только хеш, но пока и так норм
	return PasswordHash{val: val}
}

type CreatedAt struct {
	val time.Time
}

func NewCreatedAt() CreatedAt {
	return CreatedAt{val: time.Now()}
}

type UpdatedAt struct {
	val time.Time
}

type DeletedAt struct {
	val time.Time
}
