package room

import (
	"chaterley/internal/app/core"
	"chaterley/internal/app/message"
	"chaterley/internal/app/user"
	"context"
)

type RoomUseCase struct {
	roomRepo core.Repository[Room]
	userRepo core.Repository[user.User]
	msgRepo  core.Repository[message.Message]
}

func NewRoomUseCase(
	roomRepo core.Repository[Room],
	userRepo core.Repository[user.User],
	msgRepo core.Repository[message.Message],
) *RoomUseCase {
	return &RoomUseCase{roomRepo: roomRepo, userRepo: userRepo, msgRepo: msgRepo}
}

func (r *RoomUseCase) CreateRoom(
	ctx context.Context,
	name string,
	memberIDs []user.UserID,
) error {
	room, err := NewRoom(name)
	if err != nil {
		return err
	}
	if err = room.CheckMemberCount(memberIDs); err != nil {
		return err
	}
	existingUserIDs, err := r.userRepo.ExistsIds(ctx, memberIDs)
	if err != nil {
		return err
	}
	if diff := r.checkDiff(memberIDs, existingUserIDs); len(diff) > 0 {
		return core.ValidationError{Field: "memberIDs", Code: core.MemberIsNotExists}
	}
	if err = room.AddMembers(memberIDs); err != nil {
		return err
	}
	return r.saveRoom(ctx, room)
}

func (r *RoomUseCase) ChangeRoomName(
	ctx context.Context,
	roomID RoomID,
	name string,
) error {
	room, err := r.roomRepo.Get(ctx, roomID)
	if err != nil {
		return err
	}
	err = room.ChangeName(name)
	if err != nil {
		return err
	}
	return r.saveRoom(ctx, room)
}

func (r *RoomUseCase) DeleteRoom(ctx context.Context, roomID RoomID) error {
	room, err := r.roomRepo.Get(ctx, roomID)
	if err != nil {
		return err
	}
	err = room.Delete()
	if err != nil {
		return err
	}
	return r.saveRoom(ctx, room)
}

func (r *RoomUseCase) AddMemberToRoom(
	ctx context.Context,
	roomID RoomID,
	memberID user.UserID,
) error {
	room, err := r.roomRepo.Get(ctx, roomID)
	if err != nil {
		return err
	}
	isExists, err := r.userRepo.Exists(ctx, memberID)
	if err != nil {
		return err
	}
	if !isExists {
		return core.ValidationError{Field: "memberIDs", Code: core.MemberIsNotExists}
	}
	if err = room.AddMember(memberID); err != nil {
		return err
	}
	return r.saveRoom(ctx, room)
}

func (r *RoomUseCase) RemoveMemberFromRoom(
	ctx context.Context,
	roomID RoomID,
	memberID user.UserID,
) error {
	room, err := r.roomRepo.Get(ctx, roomID)
	if err != nil {
		return err
	}
	isExists, err := r.userRepo.Exists(ctx, memberID)
	if err != nil {
		return err
	}
	if !isExists {
		return core.ValidationError{Field: "memberIDs", Code: core.MemberIsNotExists}
	}
	err = room.RemoveMember(memberID)
	if err != nil {
		return err
	}
	return r.saveRoom(ctx, room)
}

func (r *RoomUseCase) checkDiff(
	memberIDs []user.UserID,
	existingMemberIDs map[user.UserID]struct{},
) []user.UserID {
	if len(existingMemberIDs) == 0 {
		return memberIDs
	}
	diff := make([]user.UserID, 0, len(memberIDs))
	for idx := range memberIDs {
		if _, ok := existingMemberIDs[memberIDs[idx]]; !ok {
			diff = append(diff, memberIDs[idx])
		}
	}
	return diff
}

func (r *RoomUseCase) saveRoom(ctx context.Context, room *Room) error {
	if err := r.roomRepo.Save(ctx, room); err != nil {
		return err
	}
	return nil
}
