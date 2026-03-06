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
	// Тело сообщения
	body core.MessageBody
	// Тип сообщения
	contentType core.ContentType
}

// NewMessage создает новый экземпляр структуры Message и возвращает указатель.
// В дальнейшем должна возвращать ошибку, если какое-то из полей невалидно.
func NewMessage(authorId core.EntityID, body string, contentType string) *Message {
	return &Message{
		id:          core.NewEntityID(),
		createdAt:   core.NewCreatedAt(),
		authorId:    authorId,
		isEdited:    core.NewIsEdited(),
		body:        core.NewMessageBody(body),
		contentType: core.NewContentType(contentType),
	}
}
