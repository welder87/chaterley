package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ErrIsEmpty = errors.New("Name is empty")
)

func TestName_NewName_WithoutError(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name string
		val  string
		want string
	}
	testCases := []testCase{
		{"1", "   test ", "test"},
		{"2", " Калина test ", "test"},
		{"2", " 34 - цвшодаырлова . ?;)() ", "34"},
	}
	for _, tc := range testCases {
		// создаем копию для параллельных подтестов
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
		name   string
		val    string
		obj    Name
		errMsg string
	}
	testCases := []testCase{
		{"1", "", Name{val: ""}, "Name is empty"},
		{"2", " ", Name{val: ""}, "Name is empty"},
		{"3", " Калина | ^^ ", Name{val: ""}, "Name is empty"},
	}
	for _, tc := range testCases {
		// создаем копию для параллельных подтестов
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			name, err := NewName(tc.val)
			assert.EqualValues(t, name, tc.obj)
			assert.Error(t, err)
			assert.ErrorContains(t, err, tc.errMsg)
		})
	}
}
