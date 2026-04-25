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
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func main() {
	val := core.NewCreatedAt[room.Room]()
	fmt.Println(val.String())
	writeConn, readConn := db.GetWriteDbCon(), db.GetReadDBCon()
	defer writeConn.Close()
	defer readConn.Close()
	roomRepo := repositories.NewRoomRepository(writeConn, readConn)
	userRepo := repositories.NewUserRepository(writeConn, readConn)
	messageRepo := repositories.NewMessageRepository(writeConn, readConn)
	roomUseCase := room.NewRoomUseCase(roomRepo, userRepo)
	msgUseCase := message.NewMessageUseCase(messageRepo)
	ctx := context.Background()
	mgr := manager.NewManager(roomUseCase, msgUseCase)
	mgr.LoadRooms(ctx)
	wsHandler := handlers.NewWebSocketHandler(mgr)
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recoverer.New())
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(compress.New())
	// используем JWT.
	// app.Use(limiter.New())
	// app.Use(pprof.New()) // мониторинг производительности
	// app.Use(etag.New())   // для кешированиtав браузере
	// app.Use(helmet.New()) // здесь XSSProtection
	// используем HttpOnly cookie, поэтому нужен и храним JWT в cookie,
	// а EncryptCookie - не нужен.
	// app.Use(csrf.New())
	// app.Use(timeout.New()) оборачивает конкретный хендлер для остановки по таймауту

	app.Get("/ws", wsHandler.Handle(ctx))
	app.Get("/", handlers.HandleIndex)
	log.Fatal(app.Listen(":3000"))
}
