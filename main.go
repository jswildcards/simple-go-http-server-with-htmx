package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DBNAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(fmt.Sprintf("Error connecting database: %s", err))
	}

	return db
}

func main() {
	LoadEnv()

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Static("/assets", "./assets")

	// POST /clicked
	app.Post("/clicked", func(c *fiber.Ctx) error {
		return c.Render("clicked", fiber.Map{
			"Now": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	// GET /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	log.Fatal(app.Listen(":3000"))
}
