package services

import (
	"context"
	"fmt"
	"time"

	"github.com/ankithans/youtube-api/pkg/database"
	"github.com/ankithans/youtube-api/pkg/models"
	"github.com/ankithans/youtube-api/pkg/utils"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// Get videos from youtube API and
// post it to postgresql database
func PostVideos() {

	// Use Youtube API to get videos
	youtubeService, err := youtube.NewService(
		context.Background(), option.WithAPIKey(utils.GoDotEnvVariable("YOUTUBE_API_KEY")),
	)
	if err != nil {
		panic(err)
	}

	// get latest videos on topic football
	tim := time.Now().Format(time.RFC3339)
	res, err := youtubeService.Search.List([]string{"id", "snippet"}).
		Q("football").
		MaxResults(10000).
		PublishedAfter(tim).
		Do()

	if err != nil {
		panic(err)
	}

	videos := []models.Video{}

	// storing the vidoes in video model slice
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
}
