package entities

import (
	"chaterley/internal/app/core"
)

// Message - Сообщение из Чата.
type Message struct {
	// id - идентификатор Сообщения
	id core.EntityID
	// content - содержание Сообщения
	content core.Content
	// authorID - идентификатор Пользователя (автора) Сообщения
	authorID core.EntityID
	// seen - флаг прочтения Сообщения
	seen core.Seen
	// createdAt - дата и время создания Сообщения
	createdAt core.CreatedAt
	// deletedAt - дата и время удаления Сообщения
	deletedAt core.DeletedAt
}

// NewMessage создает новый экземпляр структуры Message и возвращает указатель.
// В дальнейшем должна возвращать ошибку, если какое-то из полей невалидно.
func NewMessage(authorID core.EntityID, content string) *Message {
	return &Message{
		id:        core.NewEntityID(),
		createdAt: core.NewCreatedAt(),
		authorID:  authorID,
		seen:      core.NewSeen(),
		content:   core.NewContent(content),
	}
}
