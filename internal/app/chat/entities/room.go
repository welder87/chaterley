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
	// members - пользователи в Комнате.
	members map[core.EntityID]User
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
func NewRoom(name string, members ...User) (*Room, error) {
	newName, err := core.NewName(name)
	if err != nil {
		return nil, core.ValidationError{
			Field: "name",
			Code:  core.Unknown,
			Err:   err,
		}
	}
	if len(members) < minUserCount {
		return nil, core.ValidationError{Field: "members", Code: core.MinMemberCount}
	}
	membersByID := make(map[core.EntityID]User, minUserCount)
	addedMemberIds := make(map[core.EntityID]struct{}, minUserCount)
	removedUsers := make(map[core.EntityID]struct{}, minUserCount)
	for idx := range members {
		membersByID[members[idx].id] = members[idx]
		addedMemberIds[members[idx].id] = struct{}{}
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
		members:          membersByID,
		isDirty:          true,
		addedMemberIds:   addedMemberIds,
		removedMemberIds: removedUsers,
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
func (r *Room) AddMember(member User) error {
	if _, ok := r.members[member.id]; ok {
		return core.ValidationError{Field: "members", Code: core.MemberIsExists}
	}
	if len(r.members) > maxUserCount {
		return core.ValidationError{Field: "members", Code: core.MaxMemberCount}
	}
	r.members[member.id] = member
	r.addedMemberIds[member.id] = struct{}{}
	delete(r.removedMemberIds, member.id)
	r.updatedAt = core.NewUpdatedAt()
	r.isDirty = true
	r.changedFields["updatedAt"] = struct{}{}
	return nil
}

// RemoveMember - удаление члена Комнаты.
func (r *Room) RemoveMember(memberID core.EntityID) error {
	if _, ok := r.members[memberID]; !ok {
		return core.ValidationError{Field: "members", Code: core.MemberIsNotExists}
	}
	if len(r.members) <= minUserCount {
		return core.ValidationError{Field: "members", Code: core.MinMemberCount}
	}
	delete(r.members, memberID)
	r.removedMemberIds[memberID] = struct{}{}
	delete(r.addedMemberIds, memberID)
	r.updatedAt = core.NewUpdatedAt()
	r.isDirty = true
	r.changedFields["updatedAt"] = struct{}{}
	return nil
}

// Delete - удаление Комнаты.
func (r *Room) Delete() error {
	for memberID := range r.members {
		r.removedMemberIds[memberID] = struct{}{}
	}
	clear(r.addedMemberIds)
	clear(r.members)
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
func (r *Room) ToSnapshot() RoomSnapshot {
	members := make([]UserSnapshot, 0, len(r.members))
	for _, member := range r.members {
		members = append(members, member.ToSnapshot())
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
		Members:          members,
		AddedMemberIds:   addedMemberIds,
		RemovedMemberIds: removedMemberIds,
		ChangedFields:    changedFields,
	}
	if r.deletedAt != nil {
		deletedAt := r.deletedAt.Val()
		snapshot.DeletedAt = &deletedAt
	}
	r.clearChanges()
	return snapshot
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
	// Members - пользователи в Комнате.
	Members []UserSnapshot
	// AddedMemberIds - идентификаторы добавленных пользователей в Комнату.
	AddedMemberIds []core.EntityID
	// RemovedMemberIds - идентификаторы удаленных пользователей из Комнаты.
	RemovedMemberIds []core.EntityID
	// ChangedFields - наименования измененных полей
	ChangedFields []string
}
