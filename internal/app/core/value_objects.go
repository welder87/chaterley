package core

import (
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

type EntityID struct {
	val uuid.UUID
}

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
	id, err := uuid.NewV7()
	if err != nil {
		// считаем, что ситуация совсем неадекватная, и падаем
		panic(err)
	}
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
}

// Name - наименование.
type Name struct {
	val string
}

var zeroValue = Name{val: ""}

// NewName - создает наименование.
// Фактически это slug.
func NewName(val string) (Name, error) {
	if utf8.RuneCountInString(val) == 0 {
		return zeroValue, ErrNameEmpty
	}
	val = strings.TrimSpace(val)
	val = strings.ToLower(val)
	// Удаляем всё кроме английских букв в нижнем регистре, цифр и пробелов
	val = nonAlphanumericRegex.ReplaceAllString(val, "-")
	// Заменяем пробелы на дефисы
	val = strings.ReplaceAll(val, " ", "-")
	// Удаляем повторяющиеся дефисы
	val = multipleHyphensRegex.ReplaceAllString(val, "-")
	// Удаляем дефисы в начале и конце
	val = strings.Trim(val, "-")
	if len(val) == 0 {
		return zeroValue, ErrNameEmpty
	}
	if '1' <= val[0] && val[0] <= '9' {
		return zeroValue, ErrStartsWithDigit
	}
	return Name{val: val}, nil
}

// Val - геттер для получения наименования.
func (u Name) Val() string {
	return u.val
}

// Equal - сравнение наименований.
func (u Name) Equal(other Name) bool {
	return u.val == other.val
}

// String - приведение наименования к строке.
func (u Name) String() string {
	return u.val
}

type PasswordHash struct {
	val string
}

func NewPasswordHash(password string) PasswordHash {
	val := strings.TrimSpace(password)
	// тут должны быть проверки (бизнес-правила для пароля) с хешированием
	// мы не должны хранить входной val, только хеш, но пока и так норм
	return PasswordHash{val: val}
}

// CreatedAt - дата создания.
type CreatedAt struct {
	val time.Time
}

// NewCreatedAt - получение даты создания.
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
}

// UpdatedAt - дата обновления.
type UpdatedAt struct {
	val time.Time
}

// NewUpdatedAt - получение даты обновления.
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
}

// UpdatedAt - дата удаления.
type DeletedAt struct {
	val time.Time
}

// NewDeletedAt - - получение даты удаления.
func NewDeletedAt() DeletedAt {
	return DeletedAt{val: time.Now()}
}

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
}

// String - приведение даты удаления к строке.
func (u DeletedAt) String() string {
	return u.val.String()
}

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
