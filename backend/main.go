package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/tomy11/line-api/db"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	conn := db.NewConnection()
	defer conn.Close()

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Hello Word"})
	})

	app.Listen(":3000")
}
