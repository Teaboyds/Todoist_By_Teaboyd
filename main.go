package main

import (
	"log"

	"github.com/Teaboyds/Todoist_By_Teaboyd/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDB()

	app := fiber.New()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("hello patiphan")
	})

	log.Fatal(app.Listen(":9500"))
}
