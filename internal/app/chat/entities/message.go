package entities

import (
	"chaterley/internal/app/core"
)

// Message представляет сообщение из чата.
type Message struct {
	// id - Идентификатор сообщения
	id core.EntityID
	// Дата и время создания сообщения
	createdAt core.CreatedAt
	// Дата и время изменения сообщения
	updatedAt core.UpdatedAt
	// Дата и время удаления сообщения
	deletedAt core.DeletedAt
	// id - Автора сообщения.
	authorId core.EntityID
	// Флаг изменения сообщения
	isEdited core.IsEdited
	// Флаг прочтения сообщения
	isReaded core.IsReaded
	// Тело сообщения
	content core.MessageContent
	// Тип сообщения
	contentType core.ContentType
}

// NewMessage создает новый экземпляр структуры Message и возвращает указатель.
// В дальнейшем должна возвращать ошибку, если какое-то из полей невалидно.
func NewMessage(authorId core.EntityID, content string, contentType string) *Message {
	return &Message{
		id:          core.NewEntityID(),
		createdAt:   core.NewCreatedAt(),
		authorId:    authorId,
		isEdited:    core.NewIsEdited(),
		isReaded:    core.NewIsReaded(),
		content:     core.NewMessageContent(content),
		contentType: core.NewContentType(contentType),
	}
}
