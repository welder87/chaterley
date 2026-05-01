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
	// username := c.FormValue("username")
	// password := c.FormValue("password")
	// sess := session.FromContext(c)
	// sess.Set("username", username)
	// sess.Set("password", password)
	return c.Redirect().To("/rooms")
}
