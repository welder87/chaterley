package testutil

import (
	"chaterley/internal/app/message"
	"chaterley/internal/app/user"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

func NewMessageSnapshotFixture() *message.MessageSnapshot {
	return &message.MessageSnapshot{
		ID:        uuid.NewString(),
		CreatedAt: NewDateTimeFixture(),
		UpdatedAt: NewDateTimeFixture(),
		DeletedAt: nil,
		AuthorID:  uuid.NewString(),
		Seen:      gofakeit.Bool(),
		Content:   gofakeit.Comment(),
	}
}

func NewUserSnapshotFixture() *user.UserSnapshot {
	return &user.UserSnapshot{
		ID:        uuid.NewString(),
		Login:     gofakeit.Username(),
		Password:  gofakeit.Password(true, true, true, true, false, 4),
		CreatedAt: NewDateTimeFixture(),
		UpdatedAt: NewDateTimeFixture(),
		DeletedAt: nil,
	}
}

func NewDateTimeFixture() string {
	return gofakeit.Date().Format(time.RFC3339Nano)
}
