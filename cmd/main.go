package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/ankithans/youtube-api/pkg/database"
	"github.com/ankithans/youtube-api/pkg/routes"
	"github.com/ankithans/youtube-api/pkg/services"
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

	go checkVideos()

	// setup all routes
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func checkVideos() {
	for {
		services.PostVideos()
		fmt.Println(time.Now().UTC())
		time.Sleep(10 * time.Second)
	}
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/videos", routes.GetVideos)
}
