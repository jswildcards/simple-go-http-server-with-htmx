package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jswildcards/gotodo/database"
)

func DatabaseMiddleware(c *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
			"Header":      "Internal Server Error",
			"Description": err,
		})
	}

	c.Locals("db", db)
	return c.Next()
}
