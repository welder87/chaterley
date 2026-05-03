package room

import (
	"chaterley/internal/app/core"
	"chaterley/internal/app/user"
	"context"
)

type RoomUseCase struct {
	roomRepo core.ExtendedRepository[Room]
	userRepo core.Repository[user.User]
}

func NewRoomUseCase(
	roomRepo core.ExtendedRepository[Room],
	userRepo core.Repository[user.User],
) *RoomUseCase {
	return &RoomUseCase{roomRepo: roomRepo, userRepo: userRepo}
}

func (r *RoomUseCase) GetRooms(ctx context.Context) ([]*Room, error) {
	return r.roomRepo.GetAll(ctx)
}

func (r *RoomUseCase) CreateRoom(
	ctx context.Context,
	name string,
	creatorID user.UserID,
) error {
	room, err := NewRoom(name)
	if err != nil {
		return err
	}
	_, err = r.userRepo.Get(ctx, creatorID)
	if err != nil {
		return err
	}
	if err = room.AddMember(creatorID); err != nil {
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
	_, err = r.userRepo.Get(ctx, memberID)
	if err != nil {
		return err
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
	_, err = r.userRepo.Get(ctx, memberID)
	if err != nil {
		return err
	}
	err = room.RemoveMember(memberID)
	if err != nil {
		return err
	}
	return r.saveRoom(ctx, room)
}

func (r *RoomUseCase) saveRoom(ctx context.Context, room *Room) error {
	if err := r.roomRepo.Save(ctx, room); err != nil {
		return err
	}
	return nil
}
