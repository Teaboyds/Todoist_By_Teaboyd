package main

import (
	"log"

	"github.com/Teaboyds/Todoist_By_Teaboyd/database"
	"github.com/Teaboyds/Todoist_By_Teaboyd/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db := database.ConnectDB()

	app := fiber.New()

	app.Post("/register", handlers.Register(db))
	app.Post("/login", handlers.Login(db))
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("hello patiphan")
	})

	log.Fatal(app.Listen(":9500"))
}
