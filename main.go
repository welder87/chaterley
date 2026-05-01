package main

import (
	"chaterley/internal/app/manager"
	"chaterley/internal/app/message"
	"chaterley/internal/app/room"
	"chaterley/internal/handlers"
	"chaterley/internal/infrastructure/persistence/db"
	"chaterley/internal/infrastructure/persistence/repositories"
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
)

func main() {
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
	engine := html.New("./internal/handlers/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
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

	// app.Get("/", handlers.HandleIndex)
	app.Get("/*", static.New("./public"))
	app.Get("/", handlers.RedirectToRooms)
	app.Get("/login", handlers.ShowLogin)
	app.Post("/login", handlers.HandleLogin)
	// app.Get("/logout", handleLogout)
	showRooms := handlers.NewRoomsHandler(mgr)
	showRoom := handlers.NewRoomHandler(mgr)
	app.Get("/rooms/:room_id", showRoom.Handle)
	app.Get("/rooms", showRooms.Handle)
	// app.Get("/chat/:room", showChat)
	app.Get("/ws", wsHandler.Handle(ctx))
	app.Get("/404", handlers.Handle404)
	log.Fatal(app.Listen(":3000"))
}
