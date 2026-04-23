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
	authorID string,
	content string,
) (*Message, error) {
	msg, err := NewMessage(authorID, content)
	if err != nil {
		return &Message{}, err
	}
	if err := uc.msgRepo.Save(ctx, msg); err != nil {
		return &Message{}, err
	}
	return msg, nil
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
