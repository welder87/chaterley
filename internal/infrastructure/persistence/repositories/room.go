package repositories

import (
	"chaterley/internal/app/room"
	"context"
	"database/sql"
	"fmt"
)

type RoomRepository struct {
	writeDbConn *sql.DB
	readDbConn  *sql.DB
}

func NewRoomRepository(writeDbConn, readDbConn *sql.DB) *RoomRepository {
	return &RoomRepository{writeDbConn: writeDbConn, readDbConn: readDbConn}
}

func (r *RoomRepository) Save(ctx context.Context, entity *room.Room) error {
	tx, err := r.writeDbConn.BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return err
	}
	entityDTO, err := entity.ToSnapshot()
	if err != nil {
		return err
	}
	query := `
		INSERT INTO room(
			id,
			name,
			created_at,
			updated_at,
			deleted_at,
		) VALUES (
			?, ?, ?, ?, ?
		) ON CONFLICT(id)
		DO UPDATE SET
			name = excluded.name,
			updated_at = excluded.updated_at,
			deleted_at = excluded.deleted_at,
	`
	_, err = r.writeDbConn.ExecContext(ctx,
		query,
		entityDTO.ID,
		entityDTO.Name,
		entityDTO.CreatedAt,
		entityDTO.UpdatedAt,
		entityDTO.DeletedAt,
	)
	if err != nil {
		return err
	}
	if entityDTO.AddedMemberID != nil {
		query = `
			INSERT INTO room_user(
				room_id,
				user_id,
			) VALUES (
				?, ?
			) ON CONFLICT(room_id, user_id)
			DO NOTHING
		`
		_, err = r.writeDbConn.ExecContext(ctx,
			query,
			entityDTO.ID,
			entityDTO.AddedMemberID,
		)
		if err != nil {
			return err
		}
	}
	if entityDTO.RemovedMemberID != nil {
		query = `
			DELETE FROM room_user
			WHERE room_id = ? AND user_id = ?
		`
		_, err = r.writeDbConn.ExecContext(ctx,
			query,
			entityDTO.ID,
			entityDTO.RemovedMemberID,
		)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *RoomRepository) Remove(ctx context.Context, entity *room.Room) error {
	return nil
}

func (r *RoomRepository) Get(ctx context.Context, entityID room.RoomID) (*room.Room, error) {
	return nil, nil
}

func (r *RoomRepository) GetAll(ctx context.Context) ([]*room.Room, error) {
	stmt, err := r.readDbConn.PrepareContext(ctx, "SELECT * FROM room")
	rows, err := stmt.QueryContext(ctx, 0)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()
	var rooms []*room.Room
	for rows.Next() {
		var roomDTO room.RoomSnapshot
		err = rows.Scan(
			&roomDTO.ID,
			&roomDTO.Name,
			&roomDTO.CreatedAt,
			&roomDTO.UpdatedAt,
			&roomDTO.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		exRoom, err := room.NewRoomFromSnapshot(roomDTO)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, exRoom)
	}
	return rooms, nil
}
