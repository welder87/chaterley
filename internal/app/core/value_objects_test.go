package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	nameZeroValue = Name{val: ""}
)

func TestName_NewName_WithoutError(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		val  string
		want string
	}
	testCases := []testCase{
		{"Пробелы слева, справа без unicode", "   test ", "test"},
		{"unicode и пробелы в начале", " Калина test", "test"},
		{"Начало с J", "   J Калина 24 ", "j-24"},
		{"Много дефис между валидными частями", "test--------1", "test-1"},
		{"Много пробелов между валидными частями", "test      1", "test-1"},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			name, err := NewName(tc.val)
			assert.NoError(t, err)
			assert.Equal(t, name.Name(), tc.want)
		})
	}
}

func TestName_NewName_WithError(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		val  string
		err  error
	}
	testCases := []testCase{
		{"Пустая строка", "", ErrNameEmpty},
		{"Пустая строка с пробелами", " ", ErrNameEmpty},
		{"Только unicode символы и пробелы", " Калина | ^^ ", ErrNameEmpty},
		{"Начало с цифры", " 34 - цвшодаырлова . ?;)() ", ErrStartsWithDigit},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			name, err := NewName(tc.val)
			assert.EqualValues(t, name, nameZeroValue)
			assert.Error(t, err)
			assert.ErrorIs(t, err, tc.err)
		})
	}
}
