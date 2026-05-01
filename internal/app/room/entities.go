package room

import (
	"chaterley/internal/app/core"
	"chaterley/internal/app/user"
)

// Максимальное количество пользователей в Комнате.
const MaxUserCount int = 50

// Минимальное количество пользователей в Комнате.
const MinUserCount int = 1

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
	addedMemberID *user.UserID
	// removedMemberIDs - идентификатор удаленного Пользователя
	removedMemberID *user.UserID
}

// NewRoom создает новую Комнату.
// Возвращает ошибку core.ValidationError, если какое-то из полей невалидно.
func NewRoom(name string) (*Room, error) {
	newName, err := core.NewName[Room](name)
	if err != nil {
		return nil, core.ValidationError{Field: "name", Code: core.Unknown, Err: err}
	}
	return &Room{
		id:        core.NewEntityID[Room](),
		name:      newName,
		createdAt: core.NewCreatedAt[Room](),
		updatedAt: core.NewUpdatedAt[Room](),
		memberIDs: make(map[user.UserID]struct{}, MinUserCount),
	}, nil
}

func (r *Room) ID() RoomID {
	return r.id
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

// CheckMemberCount
func (r *Room) CheckMemberCount(memberIDs []user.UserID) error {
	if len(memberIDs) > MaxUserCount {
		return core.ValidationError{Field: "memberIDs", Code: core.MaxMemberCount}
	}
	return nil
}

// AddMember - добавление члена Комнаты.
func (r *Room) AddMember(memberID user.UserID) error {
	if _, ok := r.memberIDs[memberID]; ok {
		return core.ValidationError{Field: "memberIDs", Code: core.MemberIsExists}
	}
	if len(r.memberIDs) > MaxUserCount {
		return core.ValidationError{Field: "memberIDs", Code: core.MaxMemberCount}
	}
	r.memberIDs[memberID] = struct{}{}
	r.addedMemberID = &memberID
	r.updatedAt = core.NewUpdatedAt[Room]()
	return nil
}

// HasMember
func (r *Room) HasMember(memberID user.UserID) bool {
	_, ok := r.memberIDs[memberID]
	return ok
}

// RemoveMember - удаление члена Комнаты.
func (r *Room) RemoveMember(memberID user.UserID) error {
	if _, ok := r.memberIDs[memberID]; !ok {
		return core.ValidationError{Field: "memberIDs", Code: core.MemberIsNotExists}
	}
	if len(r.memberIDs) <= MinUserCount {
		return core.ValidationError{Field: "memberIDs", Code: core.MinMemberCount}
	}
	delete(r.memberIDs, memberID)
	r.removedMemberID = &memberID
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
		ID:        r.id.String(),
		Name:      r.name.Val(),
		CreatedAt: r.createdAt.String(),
		UpdatedAt: r.updatedAt.String(),
	}
	if r.deletedAt != nil {
		deletedAt := r.deletedAt.String()
		snapshot.DeletedAt = &deletedAt
	}
	if r.addedMemberID != nil {
		memberID := r.addedMemberID.String()
		snapshot.AddedMemberID = &memberID
	}
	if r.removedMemberID != nil {
		memberID := r.removedMemberID.String()
		snapshot.RemovedMemberID = &memberID
	}
	return snapshot, nil
}

// RoomSnapshot - структура данных - состояние Комнаты, на момент сериализации.
type RoomSnapshot struct {
	// ID - идентификатор Комнаты
	ID string
	// Name - наименование Комнаты
	Name string
	// CreatedAt - дата создания Комнаты
	CreatedAt string
	// UpdatedAt - дата обновления Комнаты
	UpdatedAt string
	// DeletedAt - дата удаления Комнаты
	DeletedAt *string
	// AddedMemberIDs - идентификатор добавленного пользователя в Комнату.
	AddedMemberID *string
	// RemovedMemberIDs - идентификатор удаленного пользователя из Комнаты.
	RemovedMemberID *string
}

func NewRoomFromSnapshot(snapshot RoomSnapshot) (*Room, error) {
	emptyMessage := Room{}
	messageID, err := core.NewExistsEntityID[Room](snapshot.ID)
	if err != nil {
		return &emptyMessage, err
	}
	name, err := core.NewExistsName[Room](snapshot.Name)
	if err != nil {
		return &emptyMessage, err
	}
	createdAt, err := core.NewExistsCreatedAt[Room](snapshot.CreatedAt)
	if err != nil {
		return &emptyMessage, err
	}
	updatedAt, err := core.NewExistsUpdatedAt[Room](snapshot.UpdatedAt)
	if err != nil {
		return &emptyMessage, err
	}
	var deletedAt *DeletedAt = nil
	if snapshot.DeletedAt != nil {
		newDeletedAt, err := core.NewExistsDeletedAt[Room](*(snapshot.DeletedAt))
		if err != nil {
			return &emptyMessage, err
		}
		deletedAt = &newDeletedAt
	}
	return &Room{
		id:        messageID,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}, nil
}
