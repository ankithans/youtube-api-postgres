package models

// Video defines the structure of a video
// object in the database
type Video struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Date        string `json:"Date"`
	Thumbnail   string `json:"Thumbnail"`
}
