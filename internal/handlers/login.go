package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func ShowLogin(c fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Error": nil,
	}, "layout")
}

func HandleLogin(c fiber.Ctx) error {
	return c.Redirect().To("/rooms")
}
