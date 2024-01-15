package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/redirect"
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

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/": "/tasks",
		},
		StatusCode: 301,
	}))

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

		template := "tasks/index"

		if c.QueryBool("partial") {
			return c.Render(template, fiber.Map{
				"Tasks": tasks,
			})
		}

		return c.Render(template, fiber.Map{
			"Tasks": tasks,
		}, "layout")
	})

	app.Get("/tasks/new", func(c *fiber.Ctx) error {
		template := "tasks/new"

		if c.QueryBool("partial") {
			return c.Render(template, nil)
		}

		return c.Render(template, nil, "layout")
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
