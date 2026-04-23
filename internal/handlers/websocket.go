// handlers/websocket.go
package handlers

import (
	"chaterley/internal/app/core"
	"chaterley/internal/app/manager"
	"chaterley/internal/app/room"
	"context"
	"encoding/json"
	"log"
	"os/user"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
)

type Session struct {
	c *websocket.Conn
}

func (s *Session) SendMessage(msg manager.MessageDTO) error {
	return nil
}

type WebSocketHandler struct {
	manager *manager.Manager
}

func NewWebSocketHandler(m *manager.Manager) *WebSocketHandler {
	return &WebSocketHandler{manager: m}
}

func (h *WebSocketHandler) Handle(ctx context.Context) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		for {
			var msg manager.MessageDTO
			if err := c.ReadJSON(&msg); err != nil {
				break
			}
			switch msg.Type {
			case manager.SendMessage:
				var payload manager.SentMessagePayload
				if err := json.Unmarshal(msg.Payload, &payload); err != nil {
					log.Printf("recv: %s", err)
				}
				msg, err := h.manager.SaveMessage(ctx, payload)
				if err != nil {
					log.Printf("recv: %s", err)
				}
				h.manager.Broadcast(msg)
			case manager.JoinRoom:
				var payload manager.JoinRoomMessagePayload
				if err := json.Unmarshal(msg.Payload, &payload); err != nil {
					log.Printf("recv: %s", err)
				}
				_, err := core.NewExistsEntityID[room.Room](payload.RoomID)
				if err != nil {
					return
				}
				_, err = core.NewExistsEntityID[user.User](payload.AuthorID)
				if err != nil {
					return
				}
				// h.manager.JoinRoom(roomID, userID, Session{c: c})
				// defer h.manager.LeaveRoom(userID)

			}
			// case "join":
			//     var payload JoinPayload
			//     if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			//         continue
			//     }
			//     fmt.Printf("Join rooms: %v\n", payload.Rooms)

			// case manager.Ping:
			//     var payload PingPayload
			//     if err := json.Unmarshal(msg.Payload, &payload); err != nil {
			//         continue
			//     }
			//     fmt.Printf("Ping at %d\n", payload.Time)
			// }
		}
	})
}
