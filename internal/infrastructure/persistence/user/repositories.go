package user

import (
	"chaterley/internal/app/user"
	"context"
	"database/sql"
)

type UserRepository struct {
	dbConn *sql.DB
}

func NewUserRepository(dbConn *sql.DB) *UserRepository {
	return &UserRepository{dbConn: dbConn}
}

func (r *UserRepository) Save(ctx context.Context, entity *user.User) error {
	entityDTO := entity.ToSnapshot()
	query := `
		INSERT INTO user(
			id,
			login,
			password,
			password_salt,
			created_at,
			updated_at,
			deleted_at
		)
		VALUES(
			?, ?, ?, ?, ?, ?, ?
		)
	`
	_, err := r.dbConn.ExecContext(
		ctx,
		query,
		entityDTO.ID,
		entityDTO.Login,
		entityDTO.Password,
		entityDTO.PasswordSalt,
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
		DELETE FROM USER WHERE id=?
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
		&userDTO.Login,
		&userDTO.Password,
		&userDTO.PasswordSalt,
		&userDTO.CreatedAt,
		&userDTO.UpdatedAt,
		&userDTO.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return user.NewUserFromSnapshot(userDTO)
}
