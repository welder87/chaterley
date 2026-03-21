package room

import (
	"chaterley/internal/app/auth/entities"
	"chaterley/internal/app/core"
	"chaterley/internal/app/message"
	"context"
	"errors"
)

type RoomUseCase struct {
	roomRepo *core.Repository[Room]
	userRepo *core.Repository[entities.User]
	msgRepo  *core.Repository[message.Message]
}

func NewRoomUseCase(
	roomRepo *core.Repository[Room],
	userRepo *core.Repository[entities.User],
	msgRepo *core.Repository[message.Message],
) *RoomUseCase {
	return &RoomUseCase{roomRepo: roomRepo, userRepo: userRepo, msgRepo: msgRepo}
}

func (r *RoomUseCase) CreateRoom(
	ctx context.Context,
	name string,
	memberIDs []core.EntityID,
) error {
	room, err := NewRoom(name)
	if err != nil {
		return err
	}
	isExists, err := r.userRepo.Exists(ctx, memberID)
	if err != nil {
		return err
	}
	for idx := range memberIDs {
		if err = room.AddMember(memberIDs[idx]); err != nil {
			return err
		}
	}
	return r.saveRoom(ctx, room)
}

func (r *RoomUseCase) ChangeRoomName(
	ctx context.Context,
	roomID core.EntityID,
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

func (r *RoomUseCase) DeleteRoom(ctx context.Context, roomID core.EntityID) error {
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
	roomID core.EntityID,
	memberID core.EntityID,
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
		return errors.New("not exists")
	}
	if err = room.AddMember(memberID); err != nil {
		return err
	}
	return r.saveRoom(ctx, room)
}

func (r *RoomUseCase) RemoveMemberFromRoom(
	ctx context.Context,
	roomID core.EntityID,
	memberID core.EntityID,
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
		return errors.New("not exists")
	}
	err = room.RemoveMember(memberID)
	if err != nil {
		return err
	}
	return r.saveRoom(ctx, room)
}

func (r *RoomUseCase) AddMessage(
	ctx context.Context,
	roomID core.EntityID,
	authorID core.EntityID,
	content string,
) error {
	room, err := r.roomRepo.Get(ctx, roomID)
	if err != nil {
		return err
	}
	msg, err := message.NewMessage(authorID, content)
	if err != nil {
		return err
	}
	err = room.AddMessage(msg.ID())
	if err != nil {
		return err
	}
	return r.saveRoom(ctx, room)
}

func (r *RoomUseCase) RemoveMessage(
	ctx context.Context,
	roomID core.EntityID,
	messageID core.EntityID,
) error {
	room, err := r.roomRepo.Get(ctx, roomID)
	if err != nil {
		return err
	}
	err = room.RemoveMessage(messageID)
	if err != nil {
		return err
	}
	return r.saveRoom(ctx, room)
}

func (r *RoomUseCase) GetMessages(
	ctx context.Context,
	name string,
	memberIds []core.EntityID,
) error {
	return nil
}

func (r *RoomUseCase) saveRoom(ctx context.Context, room *ent.Room) error {
	if err := r.roomRepo.Save(ctx, room); err != nil {
		return err
	}
	return nil
}
