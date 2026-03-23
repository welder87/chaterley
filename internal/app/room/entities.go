package room

import (
	"chaterley/internal/app/core"
	"chaterley/internal/app/message"
	"chaterley/internal/app/user"
	"time"
)

// Максимальное количество пользователей в Комнате.
const maxUserCount int = 100

// Минимальное количество пользователей в Комнату.
const minUserCount int = 2

type (
	RoomID    = core.EntityID[Room]
	Name      = core.Name[Room]
	CreatedAt = core.CreatedAt[Room]
	UpdatedAt = core.UpdatedAt[Room]
	DeletedAt = core.DeletedAt[Room]
)

// Room представляет Комнату Чата.
// Это агрегат.
type Room struct {
	// id - идентификатор Комнаты
	id RoomID
	// name - наименование Комнаты
	name Name
	// createdAt - дата создания Комнаты
	createdAt CreatedAt
	// updatedAt - дата обновления Комнаты
	updatedAt UpdatedAt
	// deletedAt - дата удаления Комнаты
	deletedAt *DeletedAt

	// Поля связей
	// memberIDs - идентификаторы Пользователей (предыдущее состояние)
	memberIDs map[user.UserID]struct{}
	// addedMemberIDs - идентификатор добавленного Пользователя
	addedMemberIDs []user.UserID
	// removedMemberIDs - идентификатор удаленного Пользователя
	removedMemberIDs []user.UserID
	// addedMessageID - идентификатор добавленного Сообщения
	addedMessageID *core.EntityID[message.Message]
	// removedMessageID - идентификатор удаленного Сообщения
	removedMessageID *core.EntityID[message.Message]
}

// NewRoom создает новую Комнату.
// Возвращает ошибку core.ValidationError, если какое-то из полей невалидно.
func NewRoom(name string) (*Room, error) {
	newName, err := core.NewName[Room](name)
	if err != nil {
		return nil, core.ValidationError{Field: "name", Code: core.Unknown, Err: err}
	}
	return &Room{
		id:               core.NewEntityID[Room](),
		name:             newName,
		createdAt:        core.NewCreatedAt[Room](),
		updatedAt:        core.NewUpdatedAt[Room](),
		memberIDs:        make(map[user.UserID]struct{}, 1),
		addedMemberIDs:   []user.UserID{},
		removedMemberIDs: []user.UserID{},
	}, nil
}

// ChangeName - смена наименования Комнаты.
func (r *Room) ChangeName(name string) error {
	if r.name.Val() == name {
		return core.ValidationError{Field: "name", Code: core.NameUnchanged}
	}
	newName, err := core.NewName[Room](name)
	if err != nil {
		return core.ValidationError{
			Field: "name",
			Code:  core.Unknown,
			Err:   err,
		}
	}
	if r.name == newName {
		return core.ValidationError{Field: "name", Code: core.NameUnchanged}
	}
	r.name = newName
	r.updatedAt = core.NewUpdatedAt[Room]()
	return nil
}

// AddMember - добавление члена Комнаты.
func (r *Room) AddMember(memberID user.UserID) error {
	if _, ok := r.memberIDs[memberID]; ok {
		return core.ValidationError{Field: "members", Code: core.MemberIsExists}
	}
	if len(r.memberIDs) > maxUserCount {
		return core.ValidationError{Field: "members", Code: core.MaxMemberCount}
	}
	r.addedMemberIDs = append(r.addedMemberIDs, memberID)
	r.updatedAt = core.NewUpdatedAt[Room]()
	return nil
}

// RemoveMember - удаление члена Комнаты.
func (r *Room) RemoveMember(memberID core.EntityID[user.User]) error {
	if _, ok := r.memberIDs[memberID]; !ok {
		return core.ValidationError{Field: "members", Code: core.MemberIsNotExists}
	}
	if len(r.memberIDs) <= minUserCount {
		return core.ValidationError{Field: "members", Code: core.MinMemberCount}
	}
	r.removedMemberIDs = append(r.addedMemberIDs, memberID)
	r.updatedAt = core.NewUpdatedAt[Room]()
	return nil
}

// AddMessage
func (r *Room) AddMessage(messageID core.EntityID[message.Message]) error {
	r.addedMessageID = &messageID
	r.updatedAt = core.NewUpdatedAt[Room]()
	return nil
}

func (r *Room) RemoveMessage(messageID core.EntityID[message.Message]) error {
	r.removedMessageID = &messageID
	r.updatedAt = core.NewUpdatedAt[Room]()
	return nil
}

// Delete - удаление Комнаты.
func (r *Room) Delete() error {
	deletedAt := core.NewDeletedAt[Room]()
	r.updatedAt, r.deletedAt = core.NewUpdatedAt[Room](), &deletedAt
	return nil
}

// ToSnapshot - сериализация состояния Комнаты.
func (r *Room) ToSnapshot() (RoomSnapshot, error) {
	snapshot := RoomSnapshot{
		ID:        r.id.Val().String(),
		Name:      r.name.Val(),
		CreatedAt: r.createdAt.Val(),
		UpdatedAt: r.updatedAt.Val(),
	}
	if r.deletedAt != nil {
		deletedAt := r.deletedAt.Val()
		snapshot.DeletedAt = &deletedAt
	}
	snapshot.AddedMemberIDs = uuidsToStrings(r.addedMemberIDs)
	snapshot.RemovedMemberIDs = uuidsToStrings(r.removedMemberIDs)
	if r.addedMessageID != nil {
		addedMessageID := r.addedMessageID.Val().String()
		snapshot.AddedMessageID = &addedMessageID
	}
	if r.removedMessageID != nil {
		removedMessageID := r.removedMessageID.Val().String()
		snapshot.RemovedMessageID = &removedMessageID
	}
	return snapshot, nil
}

func uuidsToStrings(memberIDs []core.EntityID[user.User]) []string {
	ids := make([]string, 0, len(memberIDs))
	for idx := range memberIDs {
		ids = append(ids, memberIDs[idx].Val().String())
	}
	return ids
}

// RoomSnapshot - структура данных - состояние Комнаты, на момент сериализации.
type RoomSnapshot struct {
	// ID - идентификатор Комнаты
	ID string
	// Name - наименование Комнаты
	Name string
	// CreatedAt - дата создания Комнаты
	CreatedAt time.Time
	// UpdatedAt - дата обновления Комнаты
	UpdatedAt time.Time
	// DeletedAt - дата удаления Комнаты
	DeletedAt *time.Time
	// AddedMemberIDs - идентификаторы добавленных пользователей в Комнату.
	AddedMemberIDs []string
	// RemovedMemberIDs - идентификаторы удаленных пользователей из Комнаты.
	RemovedMemberIDs []string
	AddedMessageID   *string
	RemovedMessageID *string
}
