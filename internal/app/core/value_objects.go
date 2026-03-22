package core

import (
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

// valueObject - структура для хранения generic значения объекта.
// используется для встраивания и ухода от boilerplate с методами структуры.
type valueObject[Value any] struct {
	val Value
}

// Val - геттер для generic значения объекта
func (vo valueObject[Value]) Val() Value {
	return vo.val
}

// EntityID - идентификатор сущности.
type EntityID[Structure any] struct {
	valueObject[uuid.UUID]
}

// NewEntityID - фабрика для генерации (по умолчанию) для идентификатора сущности.
// Принят uuid7 по следующим причинам:
// - Естественная сортировка и локальность.
// - Монотонность.
func NewEntityID[Struct any]() EntityID[Struct] {
	id, err := uuid.NewV7()
	if err != nil {
		// считаем, что ситуация совсем неадекватная, и падаем
		panic(err)
	}
	return EntityID[Struct]{valueObject[uuid.UUID]{val: id}}
}

// Name - наименование.
type Name[Struct any] struct {
	valueObject[string]
}

// NewName - создает наименование.
// Фактически это slug.
func NewName[Struct any](name string) (Name[Struct], error) {
	if utf8.RuneCountInString(name) == 0 {
		return Name[Struct]{}, ErrNameEmpty
	}
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	// Удаляем всё кроме английских букв в нижнем регистре, цифр и пробелов
	name = nonAlphanumericRegex.ReplaceAllString(name, "-")
	// Заменяем пробелы на дефисы
	name = strings.ReplaceAll(name, " ", "-")
	// Удаляем повторяющиеся дефисы
	name = multipleHyphensRegex.ReplaceAllString(name, "-")
	// Удаляем дефисы в начале и конце
	name = strings.Trim(name, "-")
	if len(name) == 0 {
		return Name[Struct]{}, ErrNameEmpty
	}
	if '1' <= name[0] && name[0] <= '9' {
		return Name[Struct]{}, ErrStartsWithDigit
	}
	return Name[Struct]{valueObject[string]{val: name}}, nil
}

type Login[Struct any] struct {
	valueObject[string]
}

func NewLogin[Struct any](login string) Login[Struct] {
	login = strings.TrimSpace(login)
	// тут должны быть проверки
	return Login[Struct]{valueObject[string]{val: login}}
}

type PasswordHash[Struct any] struct {
	valueObject[string]
}

func NewPasswordHash[Struct any](password string) PasswordHash[Struct] {
	val := strings.TrimSpace(password)
	// тут должны быть проверки (бизнес-правила для пароля) с хешированием
	// мы не должны хранить входной val, только хеш, но пока и так норм
	return PasswordHash[Struct]{valueObject[string]{val: val}}
}

// CreatedAt - дата создания.
type CreatedAt[Struct any] struct {
	valueObject[time.Time]
}

// NewCreatedAt - получение даты создания.
func NewCreatedAt[Struct any]() CreatedAt[Struct] {
	return CreatedAt[Struct]{valueObject[time.Time]{val: time.Now()}}
}

// UpdatedAt - дата обновления.
type UpdatedAt[Struct any] struct {
	valueObject[time.Time]
}

// NewUpdatedAt - получение даты обновления.
func NewUpdatedAt[Struct any]() UpdatedAt[Struct] {
	return UpdatedAt[Struct]{valueObject[time.Time]{val: time.Now()}}
}

// UpdatedAt - дата удаления.
type DeletedAt[Struct any] struct {
	valueObject[time.Time]
}

// NewDeletedAt - получение даты удаления.
func NewDeletedAt[Struct any]() DeletedAt[Struct] {
	return DeletedAt[Struct]{valueObject[time.Time]{val: time.Now()}}
}

type Seen[Struct any] struct {
	valueObject[bool]
}

func NewSeen[Struct any]() Seen[Struct] {
	return Seen[Struct]{valueObject[bool]{val: false}}
}

type Content[Struct any] struct {
	valueObject[string]
}

func NewContent[Struct any](content string) Content[Struct] {
	val := strings.TrimSpace(content)
	// тут должны быть проверки (бизнес-правила для контента сообщения)
	return Content[Struct]{valueObject[string]{val: val}}
}
