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
	deletedAt *DeletedAt
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

type MessageSnapshot struct {
	ID        string
	CreatedAt string
	DeletedAt *string
	AuthorID  string
	Seen      bool
	Content   string
}

func (m *Message) ToSnapshot() MessageSnapshot {
	snapshot := MessageSnapshot{
		ID:        m.id.String(),
		CreatedAt: m.createdAt.String(),
		AuthorID:  m.authorID.String(),
		Seen:      m.seen.Val(),
		Content:   m.content.String(),
	}
	if m.deletedAt != nil {
		deletedAt := m.deletedAt.String()
		snapshot.DeletedAt = &deletedAt
	}
	return snapshot
}

func NewMessageFromSnapshot(snapshot MessageSnapshot) (*Message, error) {
	emptyMessage := Message{}
	messageID, err := core.NewExistsEntityID[Message](snapshot.ID)
	if err != nil {
		return &emptyMessage, err
	}
	createdAt, err := core.NewExistsCreatedAt[Message](snapshot.CreatedAt)
	if err != nil {
		return &emptyMessage, err
	}
	deletedAt, err := core.NewExistsDeletedAt[Message](*snapshot.DeletedAt)
	if err != nil {
		return &emptyMessage, err
	}
	authorID, err := core.NewExistsEntityID[user.User](snapshot.AuthorID)
	if err != nil {
		return &emptyMessage, err
	}
	return &Message{
		id:        messageID,
		createdAt: createdAt,
		deletedAt: &deletedAt,
		seen:      core.NewExistsSeen[Message](snapshot.Seen),
		content:   core.NewContent[Message](snapshot.Content),
		authorID:  authorID,
	}, nil
}
