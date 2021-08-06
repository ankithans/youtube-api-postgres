package routes

import (
	"fmt"
	"math"
	"strconv"

	"github.com/ankithans/youtube-api/pkg/database"
	"github.com/ankithans/youtube-api/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetVideos(c *fiber.Ctx) error {
	var products []models.Video

	sql := "SELECT * FROM videos"

	if s := c.Query("s"); s != "" {
		sql = fmt.Sprintf("%s WHERE title LIKE '%%%s%%' OR description LIKE '%%%s%%'", sql, s, s)
	}

	if sort := c.Query("sort"); sort != "" {
		sql = fmt.Sprintf("%s ORDER BY date %s", sql, sort)
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage := 9
	var total int64

	database.DBConn.Raw(sql).Count(&total)

	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)

	database.DBConn.Raw(sql).Scan(&products)

	return c.JSON(fiber.Map{
		"data":      products,
		"total":     total,
		"page":      page,
		"last_page": math.Ceil(float64(total / int64(perPage))),
	})
}
