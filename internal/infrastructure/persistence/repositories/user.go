package repositories

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
	return nil
}

func (r *UserRepository) Remove(ctx context.Context, entity *user.User) error {
	return nil
}

func (r *UserRepository) Get(ctx context.Context, entityID user.UserID) (*user.User, error) {
	return nil, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*user.User, error) {
	return nil, nil
}

func (r *UserRepository) Exists(ctx context.Context, entityID user.UserID) (bool, error) {
	return true, nil
}

func (r *UserRepository) ExistsIds(
	ctx context.Context,
	entityIDs []user.UserID,
) (map[user.UserID]struct{}, error) {
	return nil, nil
}
