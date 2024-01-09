package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jswildcards/gotodo/middlewares"
	"github.com/jswildcards/gotodo/models"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Static("/", "./assets")

	app.Use(middlewares.DatabaseMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Get("/tasks", func(c *fiber.Ctx) error {
		db, ok := c.Locals("db").(*gorm.DB)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
				"Header":      "Internal Server Error",
				"Description": "Cannot find database instance",
			})
		}

		var tasks []models.Task
		db.Find(&tasks)

		return c.Render("tasks/index", fiber.Map{
			"Tasks": tasks,
		})
	})

	app.Get("/tasks/new", func(c *fiber.Ctx) error {
		return c.Render("tasks/new", nil)
	})

	app.Post("/tasks/create", func(c *fiber.Ctx) error {
		task := models.Task{
			Title: c.FormValue("title"),
			Description: sql.NullString{
				String: c.FormValue("description"),
				Valid:  true,
			},
			DoneAt: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
		}

		db, ok := c.Locals("db").(*gorm.DB)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
				"Header":      "Internal Server Error",
				"Description": "Cannot find database instance",
			})
		}
		db.Create(&task)

		var tasks []models.Task
		db.Find(&tasks)

		return c.Render("tasks/index", fiber.Map{
			"Tasks": tasks,
		})
	})

	app.Listen(":3000")
}
