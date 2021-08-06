package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/ankithans/youtube-api/pkg/database"
	"github.com/ankithans/youtube-api/pkg/routes"
)

func main() {

	// Establish connection to database
	database.DBConn = database.NewDatabase()

	// create new fiber app
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		msg := "Server is up and Running!"
		return c.SendString(msg)
	})

	// setup all routes
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func setupRoutes(app *fiber.App) {
	app.Post("/api/v1/youtube", routes.PostVideos)
}
