package entities

import (
	"chaterley/internal/app/core"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGroup_NewGroup_WithoutError(t *testing.T) {
	t.Parallel()
	for _, tc := range core.SuccessfulNameTestCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			group, err := NewGroup(tc.Val)
			assert.NoError(t, err)
			assert.True(t, group.name.Val() == tc.Want)
			assert.Equal(t, group.updatedAt, core.UpdatedAt{})
			assert.Equal(t, group.deletedAt, core.DeletedAt{})
			assert.WithinDuration(t, time.Now(), group.createdAt.Val(), time.Second)
		})
	}
}

func TestGroup_NewGroup_WithError(t *testing.T) {
	t.Parallel()
	for _, tc := range core.FailedNameTestCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			group, err := NewGroup(tc.Val)
			assert.Nil(t, group, nil)
			assert.Error(t, err)
			assert.ErrorIs(t, err, tc.Err)
		})
	}
}

func TestGroup_SetName_WithoutError(t *testing.T) {
	t.Parallel()
	for _, tc := range core.SuccessfulNameTestCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			name, err := core.NewName("previous_name")
			if err != nil {
				panic(err)
			}
			createdAt := core.NewCreatedAt()
			group := Group{
				id:        core.NewEntityID(),
				name:      name,
				createdAt: createdAt,
			}
			err = group.SetName(tc.Val)
			assert.NoError(t, err)
			assert.True(t, group.name.Val() == tc.Want)
			assert.Equal(t, group.createdAt, createdAt)
			assert.WithinDuration(t, time.Now(), group.updatedAt.Val(), time.Second)
			assert.Equal(t, group.deletedAt, core.DeletedAt{})
		})
	}
}

func TestGroup_SetName_WithError(t *testing.T) {
	t.Parallel()
	for _, tc := range core.FailedNameTestCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			name, err := core.NewName("previous_name")
			if err != nil {
				panic(err)
			}
			createdAt := core.NewCreatedAt()
			group := Group{
				id:        core.NewEntityID(),
				name:      name,
				createdAt: createdAt,
			}
			err = group.SetName(tc.Val)
			assert.True(t, group.name.Val() == "previous-name")
			assert.Equal(t, group.createdAt, createdAt)
			assert.Equal(t, group.updatedAt, core.UpdatedAt{})
			assert.Equal(t, group.deletedAt, core.DeletedAt{})
			assert.Error(t, err)
			assert.ErrorIs(t, err, tc.Err)
		})
	}
}

func TestGroup_Delete_WithoutError(t *testing.T) {
	t.Parallel()
	for _, tc := range core.SuccessfulNameTestCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			name, err := core.NewName("name")
			if err != nil {
				panic(err)
			}
			createdAt := core.NewCreatedAt()
			group := Group{
				id:        core.NewEntityID(),
				name:      name,
				createdAt: createdAt,
			}
			err = group.Delete()
			assert.NoError(t, err)
			assert.True(t, group.name.Val() == "name")
			assert.Equal(t, group.createdAt, createdAt)
			assert.WithinDuration(t, time.Now(), group.updatedAt.Val(), time.Second)
			assert.WithinDuration(t, time.Now(), group.deletedAt.Val(), time.Second)
		})
	}
}
