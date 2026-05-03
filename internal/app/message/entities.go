package message

import (
	"chaterley/internal/app/core"
	"chaterley/internal/app/room"
	"chaterley/internal/app/user"
)

type MessageID = core.EntityID[Message]

type (
	GroupID   = core.EntityID[Message]
	CreatedAt = core.CreatedAt[Message]
	UpdatedAt = core.UpdatedAt[Message]
	DeletedAt = core.DeletedAt[Message]
	Content   = core.Content[Message]
)

// Message - Сообщение из Чата.
type Message struct {
	// id - идентификатор Сообщения
	id MessageID
	// content - содержание Сообщения
	content Content
	// authorID - идентификатор Пользователя (автора) Сообщения
	authorID user.UserID
	// roomID - идентификатор Комнаты
	roomID room.RoomID
	// createdAt - дата и время создания Сообщения
	createdAt CreatedAt
	// updatedAt - дата и время обновления Сообщения
	updatedAt UpdatedAt
	// deletedAt - дата и время удаления Сообщения
	deletedAt *DeletedAt
}

// NewMessage создает новый экземпляр структуры Message и возвращает указатель.
// В дальнейшем должна возвращать ошибку, если какое-то из полей невалидно.
func NewMessage(roomID string, authorID string, content string) (*Message, error) {
	newContent := core.NewContent[Message](content)
	newAuthorID, err := core.NewExistsEntityID[user.User](authorID)
	if err != nil {
		return &Message{}, err
	}
	newRoomID, err := core.NewExistsEntityID[room.Room](roomID)
	if err != nil {
		return &Message{}, err
	}
	return &Message{
		id:        core.NewEntityID[Message](),
		createdAt: core.NewCreatedAt[Message](),
		updatedAt: core.NewUpdatedAt[Message](),
		authorID:  newAuthorID,
		roomID:    newRoomID,
		content:   newContent,
	}, nil
}

func (m *Message) ID() MessageID {
	return m.id
}

type MessageSnapshot struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt *string
	AuthorID  string
	RoomID    string
	Content   string
}

func (m *Message) ToSnapshot() MessageSnapshot {
	snapshot := MessageSnapshot{
		ID:        m.id.String(),
		CreatedAt: m.createdAt.String(),
		UpdatedAt: m.updatedAt.String(),
		AuthorID:  m.authorID.String(),
		RoomID:    m.roomID.String(),
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
	updatedAt, err := core.NewExistsUpdatedAt[Message](snapshot.UpdatedAt)
	if err != nil {
		return &emptyMessage, err
	}
	var deletedAt *DeletedAt = nil
	if snapshot.DeletedAt != nil {
		newDeletedAt, err := core.NewExistsDeletedAt[Message](*(snapshot.DeletedAt))
		if err != nil {
			return &emptyMessage, err
		}
		deletedAt = &newDeletedAt
	}
	authorID, err := core.NewExistsEntityID[user.User](snapshot.AuthorID)
	if err != nil {
		return &emptyMessage, err
	}
	roomID, err := core.NewExistsEntityID[room.Room](snapshot.RoomID)
	if err != nil {
		return &emptyMessage, err
	}
	return &Message{
		id:        messageID,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
		content:   core.NewContent[Message](snapshot.Content),
		authorID:  authorID,
		roomID:    roomID,
	}, nil
}
