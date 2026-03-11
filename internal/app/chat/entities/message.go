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
	// Флаг прочтения сообщения
	isReaded core.IsReaded
	// Тело сообщения
	content core.MessageContent
	// Тип сообщения
	contentType core.ContentType
}

// NewMessage создает новый экземпляр структуры Message и возвращает указатель.
// В дальнейшем должна возвращать ошибку, если какое-то из полей невалидно.
func NewMessage(authorId core.EntityID, content string, contentType string) (*Message, error) {
	newMessageContent, err := core.NewMessageContent(content)
	if err != nil {
		return nil, core.ValidationError{
			Field:  "content",
			Reason: "Is empty",
			Err:    err,
		}
	}

	newContentType, err := core.NewContentType(contentType)
	if err != nil {
		return nil, core.ValidationError{
			Field:  "contentType",
			Reason: "Is empty",
			Err:    err,
		}
	}

	return &Message{
		id:          core.NewEntityID(),
		createdAt:   core.NewCreatedAt(),
		authorId:    authorId,
		isReaded:    core.NewIsReaded(),
		content:     newMessageContent,
		contentType: newContentType,
	}, nil
}

// ID - геттер для получения идентификатора сообщения.
func (m *Message) ID() core.EntityID {
	return m.id
}

// CreatedAt - геттер для получения даты создания сообщения.
func (m *Message) CreatedAt() core.CreatedAt {
	return m.createdAt
}

// UpdatedAt - геттер для получения даты обновления сообщения.
func (m *Message) UpdatedAt() core.UpdatedAt {
	return m.updatedAt
}

// DeletedAt - геттер для получения даты удаления сообщения.
func (m *Message) DeletedAt() core.DeletedAt {
	return m.deletedAt
}

// AuthorID - геттер для получения идентификатора владельца сообщения.
func (m *Message) AuthorID() core.EntityID {
	return m.id
}

// IsReaded - геттер для получения флага о прочитанном сообщении.
func (m *Message) IsReaded() core.IsReaded {
	return m.isReaded
}

// Content - геттер для получения контента сообщения.
func (m *Message) Content() core.MessageContent {
	return m.content
}

// ContentType - геттер для получения типа контента сообщения.
func (m *Message) ContentType() core.ContentType {
	return m.contentType
}

// SetContent - сеттер для присваивания нового тела сообщения.
func (m *Message) SetContent(content string) error {
	newContent, err := core.NewMessageContent(content)
	if err != nil {
		return core.ValidationError{
			Field:  "content",
			Reason: "Is empty",
			Err:    err,
		}
	}

	m.content = newContent
	m.updatedAt = core.NewUpdatedAt()
	return nil
}

// SetContentType - сеттер для присваивания нового типа контента сообщения.
func (m *Message) SetContentType(contentType string) error {
	newContentType, err := core.NewContentType(contentType)
	if err != nil {
		return core.ValidationError{
			Field:  "contentType",
			Reason: "Is empty",
			Err:    err,
		}
	}

	m.contentType = newContentType
	m.updatedAt = core.NewUpdatedAt()
	return nil
}
