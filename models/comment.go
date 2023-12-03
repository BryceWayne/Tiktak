package models

import "time"

type Comment struct {
    ID        string    `json:"id"`
    VideoID   string    `json:"videoId"`
    UserID    string    `json:"userId"`
    Text      string    `json:"text"`
    Timestamp time.Time `json:"timestamp"`
}
