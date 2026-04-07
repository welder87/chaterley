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
	query := `
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
	snapshot := testutil.NewMessageSnapshotFixture()
	entityID, err := core.NewExistsEntityID[message.Message](snapshot.ID)
	s.Require().NoError(err)
	_, err = s.db.ExecContext(ctx,
		query,
		snapshot.ID,
		snapshot.CreatedAt,
		snapshot.UpdatedAt,
		nil,
		snapshot.AuthorID,
		snapshot.Seen,
		snapshot.Content,
	)
	s.Require().NoError(err)
	ans, err := s.repo.Get(ctx, entityID)
	s.Require().NoError(err)
	s.Equal(ans.ID(), entityID)
	s2 := ans.ToSnapshot()
	s.Equal(snapshot, &s2)
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
