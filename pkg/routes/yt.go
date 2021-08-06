package routes

import (
	"github.com/gofiber/fiber/v2"
)

func PostVideos(c *fiber.Ctx) error {

	// Use Youtube API to get videos

	// Store them in database

	return c.JSON("hi")
}
