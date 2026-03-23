package group

import (
	"chaterley/internal/app/core"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGroup_NewGroup_WithoutError(t *testing.T) {
	for _, tc := range core.SuccessfulNameTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			group, err := NewGroup(tc.Val)
			assert.NoError(t, err)
			assert.True(t, group.name.Val() == tc.Want)
			assert.WithinDuration(t, time.Now(), group.createdAt.Val(), time.Second)
			assert.WithinDuration(t, time.Now(), group.updatedAt.Val(), time.Second)
			assert.Nil(t, group.deletedAt)
		})
	}
}

func TestGroup_NewGroup_WithError(t *testing.T) {
	for _, tc := range core.FailedNameTestCases {
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
	for _, tc := range core.SuccessfulNameTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			name, err := core.NewName[Group]("previous_name")
			if err != nil {
				panic(err)
			}
			createdAt := core.NewCreatedAt[Group]()
			group := Group{
				id:        core.NewEntityID[Group](),
				name:      name,
				createdAt: createdAt,
			}
			err = group.SetName(tc.Val)
			assert.NoError(t, err)
			assert.True(t, group.name.Val() == tc.Want)
			assert.Equal(t, group.createdAt, createdAt)
			assert.WithinDuration(t, time.Now(), group.updatedAt.Val(), time.Second)
			assert.Nil(t, group.deletedAt)
		})
	}
}

func TestGroup_SetName_WithError(t *testing.T) {
	for _, tc := range core.FailedNameTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			name, err := core.NewName[Group]("previous_name")
			if err != nil {
				panic(err)
			}
			createdAt := core.NewCreatedAt[Group]()
			group := Group{
				id:        core.NewEntityID[Group](),
				name:      name,
				createdAt: createdAt,
			}
			err = group.SetName(tc.Val)
			assert.True(t, group.name.Val() == "previous-name")
			assert.Equal(t, group.createdAt, createdAt)
			assert.Equal(t, group.updatedAt, core.UpdatedAt[Group]{})
			assert.Nil(t, group.deletedAt)
			assert.Error(t, err)
			assert.ErrorIs(t, err, tc.Err)
		})
	}
}

func TestGroup_Delete_WithoutError(t *testing.T) {
	for _, tc := range core.SuccessfulNameTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			name, err := core.NewName[Group]("name")
			if err != nil {
				panic(err)
			}
			createdAt := core.NewCreatedAt[Group]()
			group := Group{
				id:        core.NewEntityID[Group](),
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
