package handlers

import (
	"chaterley/internal/app/manager"

	"github.com/gofiber/fiber/v3"
)

type RoomData struct {
	Name   string
	RoomID string
}

type RoomsHandler struct {
	manager *manager.Manager
}

func NewRoomsHandler(m *manager.Manager) *RoomsHandler {
	return &RoomsHandler{manager: m}
}

func (h *RoomsHandler) Handle(c fiber.Ctx) error {
	rooms := make([]RoomData, len(h.manager.Rooms))
	for _, room := range h.manager.Rooms {
		roomData, err := room.ToSnapshot()
		if err != nil {
			return err
		}
		rooms = append(rooms, RoomData{Name: roomData.Name, RoomID: roomData.ID})
	}
	return c.Render("rooms", fiber.Map{
		"Username": "zzz",
		"Rooms":    rooms,
	}, "layout")
}

func RedirectToRooms(c fiber.Ctx) error {
	return c.Redirect().To("/rooms")
}
