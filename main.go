package main

import (
	"chaterley/internal/app/core"
	"chaterley/internal/app/manager"
	"chaterley/internal/app/message"
	"chaterley/internal/app/room"
	"chaterley/internal/handlers"
	"chaterley/internal/infrastructure/persistence/db"
	"chaterley/internal/infrastructure/persistence/repositories"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	val := core.NewCreatedAt[room.Room]()
	fmt.Println(val.String())
	database := db.GetDBCon()
	defer database.Close()
	roomRepo := repositories.NewRoomRepository(database)
	userRepo := repositories.NewUserRepository(database)
	messageRepo := repositories.NewMessageRepository(database)
	roomUseCase := room.NewRoomUseCase(roomRepo, userRepo, messageRepo)
	msgUseCase := message.NewMessageUseCase(messageRepo)
	ctx := context.Background()
	mgr := manager.NewManager(roomUseCase, msgUseCase)
	mgr.LoadRooms(ctx)
	wsHandler := handlers.NewWebSocketHandler(mgr)
	app := fiber.New()
	app.Get("/ws", wsHandler.Handle(ctx))
	app.Get("/", handlers.HandleIndex)
	log.Fatal(app.Listen(":3000"))
}
