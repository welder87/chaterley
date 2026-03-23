package repositories

import (
	"chaterley/internal/app/chat/entities"
	"chaterley/internal/app/core"
	"context"
	"database/sql"
)

type MessageRepository struct {
	dbConn *sql.DB
}

func NewMessageRepository(dbConn *sql.DB) *MessageRepository {
	return &MessageRepository{dbConn: dbConn}
}

func (r *MessageRepository) Save(ctx context.Context, entity *entities.Message) error {
	entityDTO := entity.ToSnapshot()
	_, err := r.dbConn.ExecContext(ctx,
		"INSERT INTO message(id, created_at, updated_at, deleted_at, author_id, is_readed, content, content_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		entityDTO.ID,
		entityDTO.CreatedAt,
		entityDTO.UpdatedAt,
		entityDTO.DeletedAt,
		entityDTO.AuthorID,
		entityDTO.IsReaded,
		entityDTO.Content,
		entityDTO.ContentType,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *MessageRepository) Remove(ctx context.Context, entity *entities.Message) error {
	entityDTO := entity.ToSnapshot()
	_, err := r.dbConn.ExecContext(ctx,
		"DELETE FROM message WHERE id=?",
		entityDTO.ID,
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) Get(ctx context.Context, entityID core.EntityID) (*entities.Message, error) {
	// _ := r.dbConn.Exec(ctx, "dfasjf", entityID) -> {}
	// sto := entities.MessageDTO {...}
	// entities.Message.FromSnapshot(MessageDTO) -> entities.Message
	return nil, nil
}
