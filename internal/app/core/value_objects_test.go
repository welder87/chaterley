package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName_NewName_WithoutError(t *testing.T) {
	t.Parallel()
	for _, tc := range SuccessfulNameTestCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			name, err := NewName[emptyStruct](tc.Val)
			assert.NoError(t, err)
			assert.Equal(t, name.val, tc.Want)
		})
	}
}

func TestName_NewName_WithError(t *testing.T) {
	t.Parallel()
	for _, tc := range FailedNameTestCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			name, err := NewName[emptyStruct](tc.Val)
			assert.EqualValues(t, name, NameZeroValue)
			assert.Error(t, err)
			assert.ErrorIs(t, err, tc.Err)
		})
	}
}
