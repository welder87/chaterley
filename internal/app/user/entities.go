package user

import (
	"chaterley/internal/app/core"
)

type (
	UserID       = core.EntityID[User]
	CreatedAt    = core.CreatedAt[User]
	UpdatedAt    = core.UpdatedAt[User]
	DeletedAt    = core.DeletedAt[User]
	Login        = core.Login[User]
	PasswordHash = core.PasswordHash[User]
	PasswordSalt = core.PasswordSalt[User]
)

// User представляет пользователя чата.
// Основная сущность доменной области Chat.
type User struct {
	// id - Идентификатор пользователя
	id UserID
	// login - Логин пользователя
	login Login
	// password - Хеш пароля пользователя
	password PasswordHash
	// passwordSalt - Соль для хэширования пароля пользователя
	passwordSalt PasswordSalt
	// createdAt - Дата создания пользователя
	createdAt CreatedAt
	// updatedAt - Дата обновления пользователя
	updatedAt UpdatedAt
	// deletedAt - Дата удаления пользователя
	deletedAt *DeletedAt
}

// NewUser создает нового пользователя.
// В дальнейшем должна возвращать ошибку, если какое-то из полей невалидно.
func NewUser(login string, password string, passwordPepper []byte) (*User, error) {
	passwordSalt, _ := core.NewPasswordSalt[User]()
	hashedPassword, _ := core.NewPasswordHash[User](password, passwordSalt.Val(), passwordPepper)
	return &User{
		id:           core.NewEntityID[User](),
		login:        core.NewLogin[User](login),
		password:     hashedPassword,
		passwordSalt: passwordSalt,
		createdAt:    core.NewCreatedAt[User](),
		updatedAt:    core.NewUpdatedAt[User](),
	}, nil
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) ToSnapshot() UserSnapshot {
	snapshot := UserSnapshot{
		ID:           u.id.String(),
		Login:        u.login.Val(),
		Password:     u.password.String(),
		PasswordSalt: u.passwordSalt.String(),
		CreatedAt:    u.createdAt.String(),
		UpdatedAt:    u.updatedAt.String(),
	}
	if u.deletedAt != nil {
		deletedAt := u.deletedAt.String()
		snapshot.DeletedAt = &deletedAt
	}
	return snapshot
}

type UserSnapshot struct {
	// ID - Идентификатор пользователя
	ID string
	// Login - Логин пользователя
	Login string
	// Password - Хеш пароля пользователя
	Password string
	// PasswordSalt - Соль для хэширования пароля пользователя
	PasswordSalt string
	// CreatedAt - Дата создания пользователя
	CreatedAt string
	// UpdatedAt - Дата обновления пользователя
	UpdatedAt string
	// DeletedAt - Дата удаления пользователя
	DeletedAt *string
}

func NewUserFromSnapshot(snapshot UserSnapshot) (*User, error) {
	emptyUser := User{}
	id, err := core.NewExistsEntityID[User](snapshot.ID)
	if err != nil {
		return &emptyUser, nil
	}
	login, err := core.NewExistsLogin[User](snapshot.Login)
	if err != nil {
		return &emptyUser, err
	}
	password, err := core.NewExistsPasswordHash[User](snapshot.Password)
	if err != nil {
		return &emptyUser, err
	}
	passwordSalt, err := core.NewExistsPasswordSalt[User](snapshot.PasswordSalt)
	if err != nil {
		return &emptyUser, err
	}
	createdAt, err := core.NewExistsCreatedAt[User](snapshot.CreatedAt)
	if err != nil {
		return &emptyUser, err
	}
	updatedAt, err := core.NewExistsUpdatedAt[User](snapshot.UpdatedAt)
	if err != nil {
		return &emptyUser, err
	}

	var deletedAt *core.DeletedAt[User]
	if snapshot.DeletedAt != nil {
		val, err := core.NewExistsDeletedAt[User](*snapshot.DeletedAt)
		if err != nil {
			return &emptyUser, err
		}
		deletedAt = &val
	}
	return &User{
		id:           id,
		login:        login,
		password:     password,
		passwordSalt: passwordSalt,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
		deletedAt:    deletedAt,
	}, nil
}
