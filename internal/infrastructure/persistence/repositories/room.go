package repositories

import (
	"chaterley/internal/app/room"
	"context"
	"database/sql"
	"fmt"
)

type RoomRepository struct {
	dbConn *sql.DB
}

func NewRoomRepository(dbConn *sql.DB) *RoomRepository {
	return &RoomRepository{dbConn: dbConn}
}

func (r *RoomRepository) Save(ctx context.Context, entity *room.Room) error {
	return nil
}

func (r *RoomRepository) Remove(ctx context.Context, entity *room.Room) error {
	return nil
}

func (r *RoomRepository) Get(ctx context.Context, entityID room.RoomID) (*room.Room, error) {
	return nil, nil
}

func (r *RoomRepository) GetAll(ctx context.Context) ([]*room.Room, error) {
	stmt, err := r.dbConn.PrepareContext(ctx, "SELECT * FROM room")
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

func (r *RoomRepository) Exists(ctx context.Context, entityID room.RoomID) (bool, error) {
	return true, nil
}

func (r *RoomRepository) ExistsIds(ctx context.Context, entityIDs []room.RoomID) (map[room.RoomID]struct{}, error) {
	return nil, nil
}
