package handlers

import (
	"chaterley/internal/app/core"
	"chaterley/internal/app/manager"
	"chaterley/internal/app/message"
	"chaterley/internal/app/room"

	"github.com/gofiber/fiber/v3"
)

type RoomHandler struct {
	manager                *manager.Manager
	loadLastMessagesByRoom message.LoadLastMessagesByRoom
}

func NewRoomHandler(
	m *manager.Manager,
	loadLastMessagesByRoom message.LoadLastMessagesByRoom,
) *RoomHandler {
	return &RoomHandler{manager: m, loadLastMessagesByRoom: loadLastMessagesByRoom}
}

func (h *RoomHandler) Handle(c fiber.Ctx) error {
	roomID := c.Params("room_id")
	if roomID == "" {
		c.Redirect().To("/404")
	}
	// sess := session.FromContext(c)
	// username := sess.Get("username")
	// if username == nil {
	// 	return c.Redirect().To("/login")
	// }
	newRoomID, err := core.NewExistsEntityID[room.Room](roomID)
	if err != nil {
		c.Redirect().To("/404")
	}
	currentRoom, ok := h.manager.Rooms[newRoomID]

	if !ok {
		c.Redirect().To("/404")
	}
	roomData, err := currentRoom.ToSnapshot()
	if err != nil {
		c.Redirect().To("/404")
	}
	messages, err := h.loadLastMessagesByRoom(c.Context(), newRoomID)
	if err != nil {
		c.Redirect().To("/404")
	}
	return c.Render("room", fiber.Map{
		"Username": "zzz",
		"ID":       roomData.ID,
		"Name":     roomData.Name,
		"Messages": messages,
	}, "layout")
}
