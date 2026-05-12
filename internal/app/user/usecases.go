package user

import (
	"chaterley/internal/app/core"
	"context"
)

type UserUseCase struct {
	userRepo core.ExtendedUserRepository[User]
}

func NewUserUseCase(
	userRepo core.ExtendedUserRepository[User],
) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) CreateUser(
	login string,
	password string,
	ctx context.Context,
	secrets core.Secrets,
) (*User, error) {
	passwordPepper, err := secrets.GetPasswordPepper()
	if err != nil {
		return nil, err
	}

	us, err := NewUser(login, password, passwordPepper)
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
	secrets core.Secrets,
) (*User, error) {
	newLogin := core.NewLogin[User](login)

	usFromDB, err := uc.userRepo.FindByLogin(ctx, newLogin)
	if err != nil {
		return &User{}, err
	}

	passwordPepper, err := secrets.GetPasswordPepper()
	if err != nil {
		return nil, err
	}

	newPassword, err := core.NewPasswordHash[User](password, usFromDB.passwordSalt.Val(), passwordPepper)
	if err != nil {
		return &User{}, err
	}

	if !newPassword.Equal(usFromDB.password.Val()) {
		return &User{}, nil
	}

	us, err := NewUser(login, password, passwordPepper)
	if err != nil {
		return &User{}, err
	}

	return us, nil
}
