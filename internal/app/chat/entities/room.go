package entities

import (
	"chaterley/internal/app/core"
	"time"
)

// Максимальное количество пользователей в Комнате.
const maxUserCount int = 100

// Минимальное количество пользователей в Комнату.
const minUserCount int = 2

// Room представляет Комнату Чата.
// Это агрегат.
type Room struct {
	// id - идентификатор Комнаты
	id core.EntityID
	// name - наименование Комнаты
	name core.Name
	// memberIDs - идентификаторы Пользователей в Комнате.
	memberIDs map[core.EntityID]struct{}
	// recentMessages - недавние Сообщения в Комнате.
	recentMessages []*Message
	// lastMessageID - последнее Сообщение для пагинации.
	lastMessageID *core.EntityID
	// messages - сообщения для добавления в Комнате.
	messages []*Message
	// createdAt - дата создания Комнаты
	createdAt core.CreatedAt
	// updatedAt - дата обновления Комнаты
	updatedAt core.UpdatedAt
	// deletedAt - дата удаления Комнаты
	deletedAt *core.DeletedAt
	// isDirty - есть ли изменения, не сохраненные с хранилище.
	isDirty bool
	// addedMemberIds - идентификаторы удаленных пользователей
	addedMemberIds map[core.EntityID]struct{}
	// removedMemberIds - идентификаторы добавленных пользователей
	removedMemberIds map[core.EntityID]struct{}
	// changedFields - наименования измененных полей
	changedFields map[string]struct{}
}

// NewRoom создает новую Комнату.
// Возвращает ошибку core.ValidationError, если какое-то из полей невалидно.
func NewRoom(name string) (*Room, error) {
	newName, err := core.NewName(name)
	if err != nil {
		return nil, core.ValidationError{
			Field: "name",
			Code:  core.Unknown,
			Err:   err,
		}
	}
	changedFields := map[string]struct{}{
		"id":        {},
		"name":      {},
		"createdAt": {},
		"updatedAt": {},
	}
	return &Room{
		id:               core.NewEntityID(),
		name:             newName,
		createdAt:        core.NewCreatedAt(),
		updatedAt:        core.NewUpdatedAt(),
		memberIDs:        make(map[core.EntityID]struct{}, minUserCount),
		recentMessages:   []*Message{},
		messageCount:     0,
		isDirty:          true,
		addedMemberIds:   make(map[core.EntityID]struct{}, minUserCount),
		removedMemberIds: make(map[core.EntityID]struct{}, minUserCount),
		changedFields:    changedFields,
	}, nil
}

// ChangeName - смена наименования Комнаты.
func (r *Room) ChangeName(name string) error {
	field := "name"
	if r.name.Val() == name {
		return core.ValidationError{Field: field, Code: core.NameUnchanged}
	}
	newName, err := core.NewName(name)
	if err != nil {
		return core.ValidationError{
			Field: field,
			Code:  core.Unknown,
			Err:   err,
		}
	}
	if r.name == newName {
		return core.ValidationError{Field: field, Code: core.NameUnchanged}
	}
	r.name = newName
	r.updatedAt = core.NewUpdatedAt()
	r.isDirty = true
	r.changedFields["name"] = struct{}{}
	r.changedFields["updatedAt"] = struct{}{}
	return nil
}

// AddMember - добавление члена Комнаты.
func (r *Room) AddMember(memberId core.EntityID) error {
	if _, ok := r.memberIDs[memberId]; ok {
		return core.ValidationError{Field: "members", Code: core.MemberIsExists}
	}
	if len(r.memberIDs) > maxUserCount {
		return core.ValidationError{Field: "members", Code: core.MaxMemberCount}
	}
	r.memberIDs[memberId] = struct{}{}
	r.addedMemberIds[memberId] = struct{}{}
	delete(r.removedMemberIds, memberId)
	r.updatedAt = core.NewUpdatedAt()
	r.isDirty = true
	r.changedFields["updatedAt"] = struct{}{}
	return nil
}

// RemoveMember - удаление члена Комнаты.
func (r *Room) RemoveMember(memberID core.EntityID) error {
	if _, ok := r.memberIDs[memberID]; !ok {
		return core.ValidationError{Field: "members", Code: core.MemberIsNotExists}
	}
	if len(r.memberIDs) <= minUserCount {
		return core.ValidationError{Field: "members", Code: core.MinMemberCount}
	}
	delete(r.memberIDs, memberID)
	r.removedMemberIds[memberID] = struct{}{}
	delete(r.addedMemberIds, memberID)
	r.updatedAt = core.NewUpdatedAt()
	r.isDirty = true
	r.changedFields["updatedAt"] = struct{}{}
	return nil
}

// AddMember - добавление члена Комнаты.
func (r *Room) AddMessage(authorID core.EntityID, content string) error {
	message := NewMessage(authorID, content)
	r.messages = append(r.messages, message)
	r.memberIDs[memberId] = struct{}{}
	r.addedMemberIds[memberId] = struct{}{}
	delete(r.removedMemberIds, memberId)
	r.updatedAt = core.NewUpdatedAt()
	r.isDirty = true
	r.changedFields["updatedAt"] = struct{}{}
	return nil
}

// Delete - удаление Комнаты.
func (r *Room) Delete() error {
	for memberID := range r.memberIDs {
		r.removedMemberIds[memberID] = struct{}{}
	}
	clear(r.addedMemberIds)
	clear(r.memberIDs)
	deletedAt := core.NewDeletedAt()
	r.updatedAt, r.deletedAt = core.NewUpdatedAt(), &deletedAt
	r.isDirty = true
	r.changedFields["updatedAt"] = struct{}{}
	r.changedFields["deletedAt"] = struct{}{}
	return nil
}

func (r *Room) clearChanges() {
	r.isDirty = false
	clear(r.addedMemberIds)
	clear(r.removedMemberIds)
	clear(r.changedFields)
}

// ToSnapshot - сериализация состояния Комнаты.
func (r *Room) ToSnapshot() (RoomSnapshot, error) {
	if len(r.memberIDs) < minUserCount {
		return RoomSnapshot{}, core.ValidationError{
			Field: "members",
			Code:  core.MinMemberCount,
		}
	}
	addedMemberIds := make([]core.EntityID, 0, len(r.addedMemberIds))
	for memberId := range r.addedMemberIds {
		addedMemberIds = append(addedMemberIds, memberId)
	}
	removedMemberIds := make([]core.EntityID, 0, len(r.removedMemberIds))
	for memberId := range r.removedMemberIds {
		removedMemberIds = append(removedMemberIds, memberId)
	}
	changedFields := make([]string, 0, len(r.changedFields))
	snapshot := RoomSnapshot{
		ID:               r.id.Val(),
		Name:             r.name.Val(),
		CreatedAt:        r.createdAt.Val(),
		UpdatedAt:        r.updatedAt.Val(),
		AddedMemberIds:   addedMemberIds,
		RemovedMemberIds: removedMemberIds,
		ChangedFields:    changedFields,
	}
	if r.deletedAt != nil {
		deletedAt := r.deletedAt.Val()
		snapshot.DeletedAt = &deletedAt
	}
	r.clearChanges()
	return snapshot, nil
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
	// AddedMemberIds - идентификаторы добавленных пользователей в Комнату.
	AddedMemberIds []core.EntityID
	// RemovedMemberIds - идентификаторы удаленных пользователей из Комнаты.
	RemovedMemberIds []core.EntityID
	// ChangedFields - наименования измененных полей
	ChangedFields []string
}
