package core

import (
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

// valueObject - структура для хранения generic значения объекта.
// используется для встраивания и ухода от boilerplate с методами структуры.
type valueObject[Value any] struct {
	val Value
}

<<<<<<< HEAD
// Val - геттер для generic значения объекта
func (vo valueObject[Value]) Val() Value {
	return vo.val
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
=======
// Equal - сравнение ID.
func (e EntityID) Equal(other EntityID) bool {
	return e.val == other.val
}

// Val - геттер для получения ID.
func (e EntityID) Val() uuid.UUID {
	return e.val
}

// String - приведение ID к строке.
func (e EntityID) String() string {
	return e.val.String()
}

func NewEntityID() EntityID {
>>>>>>> 835b87d2f05fdd94f00f9e2226cac7a7b4e08d48
	id, err := uuid.NewV7()
	if err != nil {
		// считаем, что ситуация совсем неадекватная, и падаем
		panic(err)
	}
<<<<<<< HEAD
	return EntityID[Struct]{valueObject[uuid.UUID]{val: id}}
=======
	return EntityID{
		val: id,
	}
}

func NewExistsEntityID(id uuid.UUID) EntityID {
	return EntityID{
		val: id,
	}
}

type Login struct {
	val string
}

func NewLogin(val string) Login {
	val = strings.TrimSpace(val)
	// тут должны быть проверки (бизнес-правила для логина)
	return Login{val: val}
>>>>>>> 835b87d2f05fdd94f00f9e2226cac7a7b4e08d48
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

type Login[Struct any] struct {
	valueObject[string]
}

func NewLogin[Struct any](login string) Login[Struct] {
	login = strings.TrimSpace(login)
	// тут должны быть проверки
	return Login[Struct]{valueObject[string]{val: login}}
}

type PasswordHash[Struct any] struct {
	valueObject[string]
}

func NewPasswordHash[Struct any](password string) PasswordHash[Struct] {
	val := strings.TrimSpace(password)
	// тут должны быть проверки (бизнес-правила для пароля) с хешированием
	// мы не должны хранить входной val, только хеш, но пока и так норм
	return PasswordHash[Struct]{valueObject[string]{val: val}}
}

// CreatedAt - дата создания.
type CreatedAt[Struct any] struct {
	valueObject[time.Time]
}

// NewCreatedAt - получение даты создания.
<<<<<<< HEAD
func NewCreatedAt[Struct any]() CreatedAt[Struct] {
	return CreatedAt[Struct]{valueObject[time.Time]{val: time.Now()}}
=======
func NewCreatedAt() CreatedAt {
	return CreatedAt{val: time.Now()}
}

func NewExistsCreatedAt(time time.Time) CreatedAt {
	return CreatedAt{val: time}
}

// Val - геттер для получения даты создания.
func (c CreatedAt) Val() time.Time {
	return c.val
}

// Equal - сравнение даты создания с другой датой создания.
func (c CreatedAt) Equal(other CreatedAt) bool {
	return c.val.Equal(other.val)
}

// String - приведение даты создания к строке.
func (c CreatedAt) String() string {
	return c.val.String()
>>>>>>> 835b87d2f05fdd94f00f9e2226cac7a7b4e08d48
}

// UpdatedAt - дата обновления.
type UpdatedAt[Struct any] struct {
	valueObject[time.Time]
}

// NewUpdatedAt - получение даты обновления.
<<<<<<< HEAD
func NewUpdatedAt[Struct any]() UpdatedAt[Struct] {
	return UpdatedAt[Struct]{valueObject[time.Time]{val: time.Now()}}
=======
func NewUpdatedAt() UpdatedAt {
	return UpdatedAt{val: time.Now()}
}

func NewExistsUpdatedAt(time time.Time) UpdatedAt {
	return UpdatedAt{val: time}
}

// Val - геттер для получения даты обновления.
func (u UpdatedAt) Val() time.Time {
	return u.val
}

// Equal - сравнение даты обновления с другой датой обновления.
func (u UpdatedAt) Equal(other UpdatedAt) bool {
	return u.val.Equal(other.val)
}

// String - приведение даты обновления к строке.
func (u UpdatedAt) String() string {
	return u.val.String()
>>>>>>> 835b87d2f05fdd94f00f9e2226cac7a7b4e08d48
}

// UpdatedAt - дата удаления.
type DeletedAt[Struct any] struct {
	valueObject[time.Time]
}

// NewDeletedAt - получение даты удаления.
func NewDeletedAt[Struct any]() DeletedAt[Struct] {
	return DeletedAt[Struct]{valueObject[time.Time]{val: time.Now()}}
}

<<<<<<< HEAD
type Seen[Struct any] struct {
	valueObject[bool]
}

func NewSeen[Struct any]() Seen[Struct] {
	return Seen[Struct]{valueObject[bool]{val: false}}
=======
func NewExistsDeleatedAt(time time.Time) DeletedAt {
	return DeletedAt{val: time}
}

// Val - геттер для получения даты удаления.
func (u DeletedAt) Val() time.Time {
	return u.val
}

// Equal - сравнение даты удаления с другой датой удаления.
func (d DeletedAt) Equal(other DeletedAt) bool {
	return d.val.Equal(other.val)
>>>>>>> 835b87d2f05fdd94f00f9e2226cac7a7b4e08d48
}

type Content[Struct any] struct {
	valueObject[string]
}

<<<<<<< HEAD
func NewContent[Struct any](content string) Content[Struct] {
	val := strings.TrimSpace(content)
	// тут должны быть проверки (бизнес-правила для контента сообщения)
	return Content[Struct]{valueObject[string]{val: val}}
=======
type IsReaded struct {
	val bool
}

func NewIsReaded() IsReaded {
	return IsReaded{val: false}
}

func NewExistsIsReaded(val bool) IsReaded {
	return IsReaded{val: val}
}

// Equal - сравнение флагов прочитано.
func (r IsReaded) Equal(other IsReaded) bool {
	return r.val == other.val
}

// Val - геттер для получения флага прочитано.
func (r IsReaded) Val() bool {
	return r.val
}

type MessageContent struct {
	val string
}

func NewMessageContent(text string) MessageContent {
	text = strings.TrimSpace(text)
	return MessageContent{val: text}
}

// Equal - сравнение контентов.
func (c MessageContent) Equal(other MessageContent) bool {
	return c.val == other.val
}

// Val - геттер для получения контента.
func (c MessageContent) Val() string {
	return c.val
}

// String - приведение контента к строке.
func (c MessageContent) String() string {
	return c.val
}

type ContentType struct {
	val string
}

func NewContentType(val string) ContentType {
	val = strings.TrimSpace(val)
	return ContentType{val: val}
>>>>>>> 835b87d2f05fdd94f00f9e2226cac7a7b4e08d48
}

// Equal - сравнение типов контента.
func (ct ContentType) Equal(other ContentType) bool {
	return ct.val == other.val
}

// Val - геттер для получения типа контента.
func (ct ContentType) Val() string {
	return ct.val
}

// String - приведение типа контента к строке.
func (ct ContentType) String() string {
	return ct.val
}
