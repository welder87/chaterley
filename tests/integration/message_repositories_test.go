// go:build integration

package integration_test

import (
	"chaterley/internal/app/core"
	message "chaterley/internal/app/message"
	message_repo "chaterley/internal/infrastructure/persistence/message"
	"chaterley/tests/testutil"
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MessageRepositoryIntegrationSuite struct {
	BaseIntegrationSuite
	repo *message_repo.MessageRepository
}

func (s *MessageRepositoryIntegrationSuite) SetupTest() {
	s.BaseIntegrationSuite.SetupTest()
	s.repo = message_repo.NewMessageRepository(s.db)
}

func (s *MessageRepositoryIntegrationSuite) Test_Get_WithoutError() {
	ctx := context.Background()
	userSnapshot := testutil.NewUserSnapshotFixture()
	userQuery := `
		INSERT INTO user(
			id,
			login,
			password,
			created_at,
			updated_at,
			deleted_at
		)
		VALUES(
			?, ?, ?, ?, ?, ?
		)
	`
	_, err := s.db.ExecContext(ctx,
		userQuery,
		userSnapshot.ID,
		userSnapshot.Login,
		userSnapshot.Password,
		userSnapshot.CreatedAt,
		userSnapshot.UpdatedAt,
		nil,
	)
	s.Require().NoError(err)
	messageQuery := `
		INSERT INTO message(
			id,
			created_at,
			updated_at,
			deleted_at,
			author_id,
			seen,
			content
		) VALUES (
			?, ?, ?, ?, ?, ?, ?
		)
	`
	messageSnapshot := testutil.NewMessageSnapshotFixture()
	messageSnapshot.AuthorID = userSnapshot.ID
	_, err = s.db.ExecContext(ctx,
		messageQuery,
		messageSnapshot.ID,
		messageSnapshot.CreatedAt,
		messageSnapshot.UpdatedAt,
		nil,
		messageSnapshot.AuthorID,
		messageSnapshot.Seen,
		messageSnapshot.Content,
	)
	s.Require().NoError(err)
	messageID, err := core.NewExistsEntityID[message.Message](messageSnapshot.ID)
	s.Require().NoError(err)
	ans, err := s.repo.Get(ctx, messageID)
	s.Require().NoError(err)
	s.Equal(ans.ID(), messageID)
	s2 := ans.ToSnapshot()
	s.Equal(messageSnapshot, &s2)
}

func (s *MessageRepositoryIntegrationSuite) Test_Get_WithError() {
	ctx := context.Background()
	entityID := core.NewEntityID[message.Message]()
	ans, err := s.repo.Get(ctx, entityID)
	s.Require().Error(err)
	s.Require().Nil(ans)
}

func TestUserRepositoryIntegrationSuite(t *testing.T) {
	suite.Run(t, new(MessageRepositoryIntegrationSuite))
}
