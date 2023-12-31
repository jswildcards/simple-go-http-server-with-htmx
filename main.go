package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Post("/clicked", func(c *fiber.Ctx) error {
		return c.Render("clicked", fiber.Map{
			"Now": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	log.Fatal(app.Listen(":3000"))
}
