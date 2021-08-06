package models

type Video struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Date        string `json:"Date"`
	Thumbnail   string `json:"Thumbnail"`
}
