package models

import "time"

type Video struct {
    ID          string    `json:"id"`
    UserID      string    `json:"userId"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    URL         string    `json:"url"`       // URL to the video file
    Thumbnail   string    `json:"thumbnail"` // URL to the thumbnail image
    Timestamp   time.Time `json:"timestamp"`
}
