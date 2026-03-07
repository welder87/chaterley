package core

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"

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

// NewName - создает наименование группы пользователя чата.
// Фактически это slug.
func NewName(val string) (Name, error) {
	if utf8.RuneCountInString(val) == 0 {
		return Name{val: val}, errors.New("Name is empty")
	}
	val = strings.TrimSpace(val)
	val = strings.ToLower(val)
	// Удаляем всё кроме букв, цифр и пробелов
	val = nonAlphanumericRegex.ReplaceAllString(val, "-")
	// Заменяем пробелы на дефисы
	val = strings.ReplaceAll(val, " ", "-")
	// Удаляем повторяющиеся дефисы
	val = multipleHyphensRegex.ReplaceAllString(val, "-")
	// Удаляем дефисы в начале и конце
	val = strings.Trim(val, "-")
	if utf8.RuneCountInString(val) == 0 {
		return Name{val: val}, errors.New("Name is empty")
	}
	return Name{val: val}, nil
}

func (u *Name) Name() string {
	return u.val
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

type IsEdited struct {
	val bool
}

func NewIsEdited() IsEdited {
	return IsEdited{val: false}
}

type IsReaded struct {
	val bool
}

func NewIsReaded() IsReaded {
	return IsReaded{val: false}
}

type MessageContent struct {
	val string
}

func NewMessageContent(text string) MessageContent {
	text = strings.TrimSpace(text)
	return MessageContent{val: text}
}

type ContentType struct {
	val string
}

func NewContentType(val string) ContentType {
	val = strings.TrimSpace(val)
	return ContentType{val: val}
}
