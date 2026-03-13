package entities

import (
	"chaterley/internal/app/core"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRoom_NewRoom_WithoutError(t *testing.T) {
	t.Parallel()
	for _, tc := range core.SuccessfulNameTestCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			// GIVEN
			users := makeUsers(2)
			expectedUsers := make(map[core.EntityID]User, 2)
			addedUserIds := make(map[core.EntityID]struct{}, 2)
			for idx := range users {
				expectedUsers[users[idx].id] = users[idx]
				addedUserIds[users[idx].id] = struct{}{}
			}
			changedFields := map[string]struct{}{
				"id":        {},
				"name":      {},
				"createdAt": {},
				"updatedAt": {},
			}
			// WHEN
			room, err := NewRoom(tc.Val, users...)
			// THEN
			assert.NoError(t, err)
			assert.True(t, room.name.Val() == tc.Want)
			now := time.Now()
			assert.WithinDuration(t, now, room.createdAt.Val(), time.Second)
			assert.WithinDuration(t, now, room.updatedAt.Val(), time.Second)
			assert.Nil(t, room.deletedAt)
			assert.Equal(t, room.members, expectedUsers)
			assert.Equal(t, room.addedMemberIds, addedUserIds)
			assert.Equal(t, room.removedMemberIds, map[core.EntityID]struct{}{})
			assert.Equal(t, room.changedFields, changedFields)
		})
	}
}

func TestRoom_NewRoom_WithNameError(t *testing.T) {
	t.Parallel()
	for _, tc := range core.FailedNameTestCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			users := makeUsers(2)
			room, err := NewRoom(tc.Val, users...)
			assert.Nil(t, room, nil)
			assert.Error(t, err)
			assert.ErrorIs(t, err, tc.Err)
		})
	}
}

func TestRoom_NewRoom_WithMembersError(t *testing.T) {
	t.Parallel()
	type testCase struct {
		Name  string
		Users []User
		Err   error
	}
	testCases := []testCase{
		{"Без пользователей", makeUsers(0), core.ValidationError{}},
		{"Один пользователь", makeUsers(1), core.ValidationError{}},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			room, err := NewRoom("test", tc.Users...)
			assert.Nil(t, room, nil)
			assert.Error(t, err)
			assert.ErrorIs(t, err, tc.Err)
		})
	}
}

func TestRoom_AddMember_WithoutError(t *testing.T) {
	t.Parallel()
	t.Run("1", func(t *testing.T) {
		t.Parallel()
		name, _ := core.NewName("test")
		createdAt := core.NewCreatedAt()
		updatedAt := core.NewUpdatedAt()
		room := Room{
			id:        core.NewEntityID(),
			name:      name,
			createdAt: createdAt,
			updatedAt: updatedAt,
			members:   map[core.EntityID]User{},
		}
		user := User{
			id:        core.NewEntityID(),
			login:     core.NewLogin("test"),
			createdAt: createdAt,
			updatedAt: updatedAt,
		}
		err := room.AddMember(user)
		assert.NoError(t, err)
		assert.Equal(t, room.createdAt, createdAt)
		assert.Greater(t, room.updatedAt.Val(), updatedAt.Val())
		assert.WithinDuration(t, time.Now(), room.updatedAt.Val(), time.Second)
		assert.Nil(t, room.deletedAt)
		assert.Equal(t, room.members, map[core.EntityID]User{user.id: user})
	})
}

func makeUsers(cnt int) []User {
	createdAt := core.NewCreatedAt()
	updatedAt := core.NewUpdatedAt()
	return []User{
		User{
			id:        core.NewEntityID(),
			login:     core.NewLogin("test1"),
			createdAt: createdAt,
			updatedAt: updatedAt,
		},
		User{
			id:        core.NewEntityID(),
			login:     core.NewLogin("test2"),
			createdAt: createdAt,
			updatedAt: updatedAt,
		},
	}
}
