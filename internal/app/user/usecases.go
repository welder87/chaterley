package user

import (
	"chaterley/internal/app/core"
	"context"
)

type UserUseCase struct {
	userRepo core.Repository[User]
}

func NewUserUseCase(
	userRepo core.Repository[User],
) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) CreateUser(
	login string,
	password string,
	ctx context.Context,
) (*User, error) {
	us, err := NewUser(login, password)
	if err != nil {
		return &User{}, err
	}

	uc.userRepo.Save(ctx, us)
	return us, nil
}

func (uc *UserUseCase) CreateExistsUser(
	login string,
	password string,
	ctx context.Context,
) (*User, error) {
	newLogin := core.NewLogin[User](login)

	usFromDB, err := uc.userRepo.FindByLogin(ctx, newLogin)
	if err != nil {
		return &User{}, err
	}

	newPassword, err := core.NewPasswordHash[User](password, usFromDB.passwordSalt.Val())
	if err != nil {
		return &User{}, err
	}

	if newPassword != usFromDB.password {
		return &User{}, nil
	}

	us, err := NewUser(login, password)
	if err != nil {
		return &User{}, err
	}

	return us, nil
}
