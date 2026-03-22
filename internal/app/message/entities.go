package message

import (
	"chaterley/internal/app/core"
	"chaterley/internal/app/user"
)

type MessageID = core.EntityID[Message]

type (
	GroupID   = core.EntityID[Message]
	CreatedAt = core.CreatedAt[Message]
	DeletedAt = core.DeletedAt[Message]
	Content   = core.Content[Message]
	Seen      = core.Seen[Message]
)

// Message - Сообщение из Чата.
type Message struct {
	// id - идентификатор Сообщения
	id MessageID
	// content - содержание Сообщения
	content Content
	// authorID - идентификатор Пользователя (автора) Сообщения
	authorID user.UserID
	// seen - флаг прочтения Сообщения
	seen Seen
	// createdAt - дата и время создания Сообщения
	createdAt CreatedAt
	// deletedAt - дата и время удаления Сообщения
	deletedAt DeletedAt
}

// NewMessage создает новый экземпляр структуры Message и возвращает указатель.
// В дальнейшем должна возвращать ошибку, если какое-то из полей невалидно.
func NewMessage(authorID user.UserID, content string) (*Message, error) {
	newContent := core.NewContent[Message](content)
	return &Message{
		id:        core.NewEntityID[Message](),
		createdAt: core.NewCreatedAt[Message](),
		authorID:  authorID,
		seen:      core.NewSeen[Message](),
		content:   newContent,
	}, nil
}
