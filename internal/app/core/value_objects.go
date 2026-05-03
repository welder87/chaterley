package core

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"encoding/base64"

	argonize "github.com/KEINOS/go-argonize"
	"github.com/google/uuid"
)

// valueObject - структура для хранения generic значения объекта.
// используется для встраивания и ухода от boilerplate с методами структуры.
type valueObject[Value comparable] struct {
	val Value
}

type timeValueObject struct {
	valueObject[time.Time]
}

func (vo timeValueObject) String() string {
	return vo.val.Format(time.RFC3339Nano)
}

func (vo valueObject[Value]) Val() Value {
	return vo.val
}

// Equal - сравнение двух объектов одного типа по значению.
func (vo valueObject[Value]) Equal(other valueObject[Value]) bool {
	return vo.val == other.val
}

// String - приведение значения к строке.
func (vo valueObject[Value]) String() string {
	return fmt.Sprintf("%v", vo.val)
}

// EntityID - идентификатор сущности.
type EntityID[Structure any] struct {
	valueObject[uuid.UUID]
}

// NewEntityID - фабрика для генерации (по умолчанию) для идентификатора сущности.
// Принят uuid7 по следующим причинам:
// - Естественная сортировка и локальность.
// - Монотонность.
func NewEntityID[Struct any]() EntityID[Struct] {
	id, err := uuid.NewV7()
	if err != nil {
		// считаем, что ситуация совсем неадекватная, и падаем
		panic(err)
	}
	return EntityID[Struct]{valueObject[uuid.UUID]{val: id}}
}

func NewExistsEntityID[Struct any](entityID string) (EntityID[Struct], error) {
	entityUUID, err := uuid.Parse(entityID)
	if err != nil {
		return EntityID[Struct]{valueObject[uuid.UUID]{}}, err
	}
	return EntityID[Struct]{valueObject[uuid.UUID]{val: entityUUID}}, nil
}

// Name - наименование.
type Name[Struct any] struct {
	valueObject[string]
}

// NewName - создает наименование.
// Фактически это slug.
func NewName[Struct any](name string) (Name[Struct], error) {
	if utf8.RuneCountInString(name) == 0 {
		return Name[Struct]{}, ErrNameEmpty
	}
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	// Удаляем всё кроме английских букв в нижнем регистре, цифр и пробелов
	name = nonAlphanumericRegex.ReplaceAllString(name, "-")
	// Заменяем пробелы на дефисы
	name = strings.ReplaceAll(name, " ", "-")
	// Удаляем повторяющиеся дефисы
	name = multipleHyphensRegex.ReplaceAllString(name, "-")
	// Удаляем дефисы в начале и конце
	name = strings.Trim(name, "-")
	if len(name) == 0 {
		return Name[Struct]{}, ErrNameEmpty
	}
	if '1' <= name[0] && name[0] <= '9' {
		return Name[Struct]{}, ErrStartsWithDigit
	}
	return Name[Struct]{valueObject[string]{val: name}}, nil
}

func NewExistsName[Struct any](name string) (Name[Struct], error) {
	return Name[Struct]{valueObject[string]{val: name}}, nil
}

type Login[Struct any] struct {
	valueObject[string]
}

func NewLogin[Struct any](login string) Login[Struct] {
	login = strings.TrimSpace(login)
	// тут должны быть проверки
	return Login[Struct]{valueObject[string]{val: login}}
}

func NewExistsLogin[Struct any](login string) (Login[Struct], error) {
	return Login[Struct]{valueObject[string]{val: login}}, nil
}

type PasswordHash[Struct any] struct {
	valueObject[string]
}

func (ph PasswordHash[any]) String() string {
	return base64.RawStdEncoding.EncodeToString([]byte(ph.val))
}

func NewPasswordHash[Struct any](password string, salt argonize.Salt) (PasswordHash[Struct], error) {
	password = strings.TrimSpace(password)
	if len(password) < 8 {
		return PasswordHash[Struct]{}, ErrPasswordLength
	}

	bytePassword := []byte(password)
	params := argonize.NewParams()

	paper := []byte(os.Getenv("PASSWORD_PEPER"))
	salt.AddPepper(paper)
	hashedObj := argonize.HashCustom(bytePassword, salt, params)
	if !hashedObj.IsValidPassword(bytePassword) {
		return PasswordHash[Struct]{}, ErrGenPasswordHashed
	}

	return PasswordHash[Struct]{valueObject[string]{val: string(hashedObj.Hash)}}, nil
}

func NewExistsPasswordHash[Struct any](password string) (PasswordHash[Struct], error) {
	passwordFromB64, err := base64.RawStdEncoding.DecodeString(password)
	if err != nil {
		return PasswordHash[Struct]{}, ErrDecodePasswordFromB64
	}
	return PasswordHash[Struct]{valueObject[string]{val: string(passwordFromB64)}}, nil
}

// Соль для хэширования пароля
type PasswordSalt[Struct any] struct {
	val argonize.Salt
}

func (ps PasswordSalt[any]) String() string {
	return base64.RawStdEncoding.EncodeToString(ps.val)
}

func (ps PasswordSalt[any]) Val() argonize.Salt {
	return ps.val
}

func NewPasswordSalt[Struct any]() (PasswordSalt[Struct], error) {
	params := argonize.NewParams()
	salt, err := argonize.NewSalt(params.SaltLength)
	if err != nil {
		return PasswordSalt[Struct]{}, ErrGenPasswordSalt
	}

	return PasswordSalt[Struct]{val: salt}, nil
}

func NewExistsPasswordSalt[Struct any](passwordSalt string) (PasswordSalt[Struct], error) {
	passwordSaltFromB64, err := base64.RawStdEncoding.DecodeString(passwordSalt)
	if err != nil {
		return PasswordSalt[Struct]{}, ErrDecodePasswordFromB64
	}
	return PasswordSalt[Struct]{val: passwordSaltFromB64}, nil
}

// CreatedAt - дата создания.
type CreatedAt[Struct any] struct {
	timeValueObject
}

// NewCreatedAt - получение даты создания.
func NewCreatedAt[Struct any]() CreatedAt[Struct] {
	return CreatedAt[Struct]{
		timeValueObject: timeValueObject{
			valueObject[time.Time]{val: time.Now()},
		},
	}
}

func NewExistsCreatedAt[Struct any](val string) (CreatedAt[Struct], error) {
	date_time, err := time.Parse(time.RFC3339Nano, val)
	if err != nil {
		return CreatedAt[Struct]{
			timeValueObject: timeValueObject{
				valueObject[time.Time]{},
			}}, err
	}
	return CreatedAt[Struct]{
		timeValueObject: timeValueObject{
			valueObject[time.Time]{val: date_time},
		}}, nil
}

// UpdatedAt - дата обновления.
type UpdatedAt[Struct any] struct {
	timeValueObject
}

// NewUpdatedAt - получение даты обновления.
func NewUpdatedAt[Struct any]() UpdatedAt[Struct] {
	return UpdatedAt[Struct]{
		timeValueObject: timeValueObject{
			valueObject[time.Time]{val: time.Now()},
		}}
}

func NewExistsUpdatedAt[Struct any](val string) (UpdatedAt[Struct], error) {
	date_time, err := time.Parse(time.RFC3339Nano, val)
	if err != nil {
		return UpdatedAt[Struct]{
			timeValueObject: timeValueObject{
				valueObject[time.Time]{}}}, err
	}
	return UpdatedAt[Struct]{
		timeValueObject: timeValueObject{
			valueObject[time.Time]{val: date_time},
		}}, nil
}

// DeletedAt - дата удаления.
type DeletedAt[Struct any] struct {
	timeValueObject
}

// NewDeletedAt - получение даты удаления.
func NewDeletedAt[Struct any]() DeletedAt[Struct] {
	return DeletedAt[Struct]{
		timeValueObject: timeValueObject{
			valueObject[time.Time]{val: time.Now()},
		}}
}

func NewExistsDeletedAt[Struct any](val string) (DeletedAt[Struct], error) {
	date_time, err := time.Parse(time.RFC3339Nano, val)
	if err != nil {
		return DeletedAt[Struct]{
			timeValueObject: timeValueObject{
				valueObject[time.Time]{}}}, err
	}
	return DeletedAt[Struct]{
		timeValueObject: timeValueObject{
			valueObject[time.Time]{val: date_time}}}, nil
}

type Seen[Struct any] struct {
	valueObject[bool]
}

func NewSeen[Struct any]() Seen[Struct] {
	return Seen[Struct]{valueObject[bool]{val: false}}
}

func NewExistsSeen[Struct any](val bool) Seen[Struct] {
	return Seen[Struct]{valueObject[bool]{val: val}}
}

type Content[Struct any] struct {
	valueObject[string]
}

func NewContent[Struct any](content string) Content[Struct] {
	val := strings.TrimSpace(content)
	// тут должны быть проверки (бизнес-правила для контента сообщения)
	return Content[Struct]{valueObject[string]{val: val}}
}

func NewExistsContent[Struct any](val string) Content[Struct] {
	return Content[Struct]{valueObject[string]{val: val}}
}
