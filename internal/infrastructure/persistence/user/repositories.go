package user

import (
	"chaterley/internal/app/user"
	"context"
	"database/sql"
)

type UserRepository struct {
	dbConn *sql.DB
}

func (r *UserRepository) Save(ctx context.Context, entity *user.User) error {
	entityDTO := entity.ToSnapshot()
	query := `
		INSERT INTO user(
			user_id,
			login,
			password,
			created_at,
			updated_at,
			deleted_at,
		)
		VALUES(
			?, ?, ?, ?, ?, ?
		)
	`
	_, err := r.dbConn.ExecContext(
		ctx,
		query,
		entityDTO.ID,
		entityDTO.Login,
		entityDTO.Password,
		entityDTO.CreatedAt,
		entityDTO.UpdatedAt,
		entityDTO.DeletedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Remove(ctx context.Context, entity *user.User) error {
	entityDTO := entity.ToSnapshot()
	query := `
		DELETE FROM USER WHERE user_id=?
	`
	_, err := r.dbConn.ExecContext(
		ctx,
		query,
		entityDTO.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Get(ctx context.Context, entityID user.UserID) (*user.User, error) {
	userFromDB := r.dbConn.QueryRowContext(ctx,
		"SELECT * FROM user WHERE id=?",
		entityID.String(),
	)

	var userDTO user.UserSnapshot
	err := userFromDB.Scan(
		&userDTO.ID,
		&userDTO.CreatedAt,
		&userDTO.UpdatedAt,
		&userDTO.DeletedAt,
		&userDTO.Login,
		&userDTO.Password,
	)
	if err != nil {
		return nil, err
	}
	return user.NewUserFromSnapshot(userDTO)
}
