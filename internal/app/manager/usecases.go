package manager

import (
	"chaterley/internal/app/message"
	"chaterley/internal/app/room"
	"chaterley/internal/app/user"
	"context"
	"errors"
	"sync"
)

type Manager struct {
	roomUseCase *room.RoomUseCase
	msgUseCase  *message.MessageUseCase
	rooms       map[room.RoomID]*room.Room
	connections map[user.UserID]Session
	mu          sync.RWMutex
}

func NewManager(roomUseCase *room.RoomUseCase, msgUseCase *message.MessageUseCase) *Manager {
	return &Manager{
		roomUseCase: roomUseCase,
		msgUseCase:  msgUseCase,
		rooms:       make(map[room.RoomID]*room.Room, 2),
		connections: make(map[user.UserID]Session, room.MinUserCount),
	}
}

func (m *Manager) LoadRooms(ctx context.Context) error {
	rooms, err := m.roomUseCase.GetRooms(ctx)
	if err != nil {
		return err
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, room := range rooms {
		m.rooms[room.ID()] = room
	}
	return nil
}

func (m *Manager) JoinRoom(roomID room.RoomID, userID user.UserID, session Session) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	cur, ok := m.rooms[roomID]
	if !ok {
		return errors.New("No room")
	}
	if !cur.HasMember(userID) {
		return errors.New("No user in room")
	}
	if _, ok := m.connections[userID]; !ok {
		m.connections[userID] = session
	}
	return nil
}

func (m *Manager) LeaveRoom(userID user.UserID) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.connections, userID)
}

func (m *Manager) SaveMessage(
	ctx context.Context,
	payload SentMessagePayload,
) (*message.Message, error) {
	return m.msgUseCase.Create(ctx, payload.AuthorID, payload.Content)
}

func (m *Manager) Broadcast(msg *message.Message) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	payload := msg.ToSnapshot()
	for userID, conn := range m.connections {
		if userID.String() != payload.AuthorID {
			conn.SendMessage(payload)
		}
	}
}
