package core

type successfulNameTestCase struct {
	Name string
	Val  string
	Want string
}
type failedNameTestCase struct {
	Name string
	Val  string
	Err  error
}

type emptyStruct struct{}

var (
	NameZeroValue = Name[emptyStruct]{}
)

var (
	SuccessfulNameTestCases = []successfulNameTestCase{
		{"Пробелы слева, справа без unicode", "   test ", "test"},
		{"unicode и пробелы в начале", " Калина test", "test"},
		{"Начало с J", "   J Калина 24 ", "j-24"},
		{"Много дефис между валидными частями", "test--------1", "test-1"},
		{"Много пробелов между валидными частями", "test      1", "test-1"},
	}
	FailedNameTestCases = []failedNameTestCase{
		{"Пустая строка", "", ErrNameEmpty},
		{"Пустая строка с пробелами", " ", ErrNameEmpty},
		{"Только unicode символы и пробелы", " Калина | ^^ ", ErrNameEmpty},
		{"Начало с цифры", " 34 - цвшодаырлова . ?;)() ", ErrStartsWithDigit},
	}
)
