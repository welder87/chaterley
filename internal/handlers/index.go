package handlers

import "github.com/gofiber/fiber/v3"

func HandleIndex(c fiber.Ctx) error {
	return c.SendFile("index.html")
}
