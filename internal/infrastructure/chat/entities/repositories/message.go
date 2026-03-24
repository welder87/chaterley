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
	messageFromDB := r.dbConn.QueryRowContext(ctx,
		"SELECT * FROM message WHERE id=?",
		entityID.String(),
	)

	var messageDTO entities.MessageDTO
	err := messageFromDB.Scan(
		&messageDTO.ID, &messageDTO.CreatedAt, &messageDTO.UpdatedAt, &messageDTO.DeletedAt,
		&messageDTO.AuthorID, &messageDTO.IsReaded, &messageDTO.ContentType, &messageDTO.Content,
	)
	if err != nil {
		return nil, err
	}

	return entities.NewMessageFromSnapshot(messageDTO), nil
}
