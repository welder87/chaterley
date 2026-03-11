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

// Val - геттер получения ID сущности.
func (e EntityID) Val() uuid.UUID {
	return e.val
}

// Equal - сравнение текущего ID сущности с переданным.
func (e EntityID) Equal(other EntityID) bool {
	return e.val == other.val
}

// String - приведение ID сущности к строке.
func (e EntityID) String() string {
	return e.val.String()
}

type Login struct {
	val string
}

func NewLogin(val string) Login {
	val = strings.TrimSpace(val)
	// тут должны быть проверки (бизнес-правила для логина)
	return Login{val: val}
}

// Val - геттер получения логина.
func (l Login) Val() string {
	return l.val
}

// Equal - сравнение текущего логина с переданным.
func (l Login) Equal(other Login) bool {
	return l.val == other.val
}

// String - приведение логина к строке.
func (l Login) String() string {
	return l.val
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

// Val - геттер для получения даты создания.
func (c CreatedAt) Val() time.Time {
	return c.val
}

// Equal - сравнение даты создания с другой датой создания.
func (c CreatedAt) Equal(other CreatedAt) bool {
	return c.val == other.val
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

// Val - геттер для получения даты обновления.
func (u UpdatedAt) Val() time.Time {
	return u.val
}

// Equal - сравнение даты обновления с другой датой обновления.
func (u UpdatedAt) Equal(other UpdatedAt) bool {
	return u.val == other.val
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

// Val - геттер для получения даты удаления.
func (u DeletedAt) Val() time.Time {
	return u.val
}

// Equal - сравнение даты удаления с другой датой удаления.
func (u DeletedAt) Equal(other DeletedAt) bool {
	return u.val == other.val
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

// Val - геттер получения флага о прочтении сообщения.
func (r IsReaded) Val() bool {
	return r.val
}

// Equal - сравнение текущего флага с переданным.
func (r IsReaded) Equal(other IsReaded) bool {
	return r.val == other.val
}

type MessageContent struct {
	val string
}

var emptyMessageContent = MessageContent{val: ""}

func NewMessageContent(text string) (MessageContent, error) {
	text = strings.TrimSpace(text)
	// Недопускается отправка пустого сообщения
	if text == "" {
		return emptyMessageContent, ErrContentEmpty
	}

	return MessageContent{val: text}, nil
}

// Val - геттер получения контента сообщения.
func (m MessageContent) Val() string {
	return m.val
}

// Equal - сравнение текущего сообщения с переданным.
func (m MessageContent) Equal(other MessageContent) bool {
	return m.val == other.val
}

// String - получение строкого представления сообщения.
func (m MessageContent) String() string {
	return m.val
}

type ContentType struct {
	val string
}

var emptyContentType = ContentType{val: ""}

func NewContentType(val string) (ContentType, error) {
	val = strings.TrimSpace(val)
	// Недопускается наличие неопределенного или пустого контента
	if val == "" {
		return emptyContentType, ErrContentTypeEmpty
	}

	return ContentType{val: val}, nil
}

// Val - геттер получение типа контента.
func (c ContentType) Val() string {
	return c.val
}

// Equal - сравнение текущего типа контента с переданным.
func (c ContentType) Equal(other ContentType) bool {
	return c.val == other.val
}

// String - получение строкого представления типа контента.
func (c ContentType) String() string {
	return c.val
}
