package room

import (
	"chaterley/internal/app/core"
	"chaterley/internal/app/user"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRoom_NewRoom_WithoutError(t *testing.T) {
	for _, tc := range core.SuccessfulNameTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			room, err := NewRoom(tc.Val)
			assert.NoError(t, err)
			assert.True(t, room.name.Val() == tc.Want)
			now := time.Now()
			assert.WithinDuration(t, now, room.createdAt.Val(), time.Second)
			assert.WithinDuration(t, now, room.updatedAt.Val(), time.Second)
			assert.Nil(t, room.deletedAt)
			assert.Equal(t, room.memberIDs, map[user.UserID]struct{}{})
			assert.Equal(t, room.addedMemberIDs, []user.UserID{})
			assert.Equal(t, room.removedMemberIDs, []user.UserID{})
		})
	}
}

func TestRoom_NewRoom_WithNameError(t *testing.T) {
	for _, tc := range core.FailedNameTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			room, err := NewRoom(tc.Val)
			assert.Nil(t, room, nil)
			assert.Error(t, err)
			assert.ErrorIs(t, err, tc.Err)
		})
	}
}

func TestRoom_AddMember_WithoutError(t *testing.T) {
	name, _ := core.NewName[Room]("test")
	createdAt := core.NewCreatedAt[Room]()
	updatedAt := core.NewUpdatedAt[Room]()
	room := Room{
		id:        core.NewEntityID[Room](),
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
		memberIDs: map[user.UserID]struct{}{},
	}
	userID := core.NewEntityID[user.User]()
	err := room.AddMember(userID)
	assert.NoError(t, err)
	assert.Equal(t, room.createdAt, createdAt)
	assert.Greater(t, room.updatedAt.Val(), updatedAt.Val())
	assert.WithinDuration(t, time.Now(), room.updatedAt.Val(), time.Second)
	assert.Nil(t, room.deletedAt)
	assert.Equal(t, room.memberIDs, map[user.UserID]struct{}{})
}
