package repositories

import (
	"chaterley/internal/app/message"
	"context"
	"database/sql"
)

type MessageRepository struct {
	writeDbConn *sql.DB
	readDbConn  *sql.DB
}

func NewMessageRepository(writeDbConn, readDbConn *sql.DB) *MessageRepository {
	return &MessageRepository{writeDbConn: writeDbConn, readDbConn: readDbConn}
}

func (r *MessageRepository) Save(ctx context.Context, entity *message.Message) error {
	entityDTO := entity.ToSnapshot()
	query := `
		INSERT INTO message(
			id,
			created_at,
			updated_at,
			deleted_at,
			author_id,
			room_id,
			content
		) VALUES (
			?, ?, ?, ?, ?, ?, ?
		)
	`
	_, err := r.writeDbConn.ExecContext(ctx,
		query,
		entityDTO.ID,
		entityDTO.CreatedAt,
		entityDTO.UpdatedAt,
		entityDTO.DeletedAt,
		entityDTO.AuthorID,
		entityDTO.RoomID,
		entityDTO.Content,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *MessageRepository) Remove(ctx context.Context, entity *message.Message) error {
	entityDTO := entity.ToSnapshot()
	_, err := r.writeDbConn.ExecContext(
		ctx,
		"DELETE FROM message WHERE id=?",
		entityDTO.ID,
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) Get(ctx context.Context, entityID message.MessageID) (*message.Message, error) {
	messageFromDB := r.readDbConn.QueryRowContext(ctx,
		"SELECT * FROM message WHERE id=?",
		entityID.String(),
	)

	var messageDTO message.MessageSnapshot
	err := messageFromDB.Scan(
		&messageDTO.ID,
		&messageDTO.CreatedAt,
		&messageDTO.UpdatedAt,
		&messageDTO.DeletedAt,
		&messageDTO.AuthorID,
		&messageDTO.Content,
	)
	if err != nil {
		return nil, err
	}
	return message.NewMessageFromSnapshot(messageDTO)
}
