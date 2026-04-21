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
	rawPassword, err := core.NewPassword[User](password)
	if err != nil {
		return &User{}, err
	}

	us, err := NewUser(login, rawPassword.String())
	if err != nil {
		return &User{}, err
	}

	uc.userRepo.Save(ctx, us)
	return us, nil
}
