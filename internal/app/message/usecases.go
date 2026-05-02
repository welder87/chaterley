package message

import (
	"chaterley/internal/app/core"
	"context"
)

type MessageUseCase struct {
	msgRepo core.Repository[Message]
}

func NewMessageUseCase(msgRepo core.Repository[Message]) *MessageUseCase {
	return &MessageUseCase{msgRepo: msgRepo}
}

func (uc *MessageUseCase) Create(
	ctx context.Context,
	msg MessageDTO,
) (*Message, error) {
	newMsg, err := NewMessage(msg.RoomID, msg.AuthorID, msg.Content)
	if err != nil {
		return newMsg, err
	}
	if err := uc.msgRepo.Save(ctx, newMsg); err != nil {
		return newMsg, err
	}
	return newMsg, nil
}

func (uc *MessageUseCase) Delete(ctx context.Context, msgID MessageID) error {
	msg, err := uc.msgRepo.Get(ctx, msgID)
	if err != nil {
		return err
	}
	if err = uc.msgRepo.Remove(ctx, msg); err != nil {
		return err
	}
	return nil
}

type MessageDTO struct {
	RoomID   string
	AuthorID string
	Content  string
}
