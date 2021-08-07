package routes

import (
	"fmt"
	"math"
	"strconv"

	"github.com/ankithans/youtube-api/pkg/database"
	"github.com/ankithans/youtube-api/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// GetVideos returns the videos from database
// according the queries provided by the user
func GetVideos(c *fiber.Ctx) error {
	var products []models.Video

	// basic sql query
	sql := "SELECT * FROM videos"

	// if your searches for video using title ot description
	if s := c.Query("s"); s != "" {
		sql = fmt.Sprintf("%s WHERE title LIKE '%%%s%%' OR description LIKE '%%%s%%'", sql, s, s)
	}

	// sort by asc or desc
	if sort := c.Query("sort"); sort != "" {
		sql = fmt.Sprintf("%s ORDER BY date %s", sql, sort)
	}

	// page number specified by user
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage := 9
	var total int64

	// total videos
	database.DBConn.Raw(sql).Count(&total)

	// sql query to get videos on particular page
	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)
	database.DBConn.Raw(sql).Scan(&products)

	// return the result to user
	return c.JSON(fiber.Map{
		"data":      products,
		"total":     total,
		"page":      page,
		"last_page": math.Ceil(float64(total / int64(perPage))),
	})
}
