package handlers

import "github.com/gofiber/fiber/v3"

func Handle404(c fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendString("404 Page Not Found")
}
