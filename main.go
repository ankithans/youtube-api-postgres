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

	// enable cors
	app.Use(cors.New())

	// root route
	app.Get("/", func(c *fiber.Ctx) error {
		msg := "Server is up and Running!"
		return c.SendString(msg)
	})

	// check new videos every 10 second
	// and push them to database
	go checkVideos()

	// setup all routes
	setupRoutes(app)

	// listen to port 3000
	log.Fatal(app.Listen(":3000"))
}

// check new videos every 10 second
// and push them to database
func checkVideos() {
	for {
		services.PostVideos()
		fmt.Println(time.Now().UTC())
		time.Sleep(10 * time.Second)
	}
}

// Setups routes
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/videos", routes.GetVideos)
}
