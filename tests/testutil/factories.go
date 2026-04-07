package testutil

import (
	"chaterley/internal/app/message"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

func NewMessageSnapshotFixture() *message.MessageSnapshot {
	return &message.MessageSnapshot{
		ID:        uuid.NewString(),
		CreatedAt: gofakeit.Date().UTC().Format(time.RFC3339Nano),
		UpdatedAt: gofakeit.Date().UTC().Format(time.RFC3339Nano),
		DeletedAt: nil,
		AuthorID:  uuid.NewString(),
		Seen:      gofakeit.Bool(),
		Content:   gofakeit.Comment(),
	}
}
