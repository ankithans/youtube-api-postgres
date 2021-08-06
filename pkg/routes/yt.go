package routes

import (
	"context"
	"fmt"

	"github.com/ankithans/youtube-api/pkg/database"
	"github.com/ankithans/youtube-api/pkg/models"
	"github.com/ankithans/youtube-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func PostVideos(c *fiber.Ctx) error {

	// Use Youtube API to get videos
	youtubeService, err := youtube.NewService(
		context.Background(), option.WithAPIKey(utils.GoDotEnvVariable("YOUTUBE_API_KEY")),
	)
	if err != nil {
		panic(err)
	}

	// tim := time.Now().Format(time.RFC3339)
	res, err := youtubeService.Search.List([]string{"id", "snippet"}).
		Q("football").
		MaxResults(10000).
		// PublishedAfter(tim).
		Do()

	if err != nil {
		panic(err)
	}

	videos := []models.Video{}

	for _, item := range res.Items {
		fmt.Println(item)
		if item.Id.Kind == "youtube#video" {
			vid := models.Video{
				Title:       item.Snippet.Title,
				Description: item.Snippet.Description,
				Date:        item.Snippet.PublishedAt,
				Thumbnail:   item.Snippet.Thumbnails.Default.Url,
			}
			videos = append(videos, vid)
		}
	}

	// Store them in database
	database.DBConn.Create(&videos)

	return c.JSON(videos)
}
