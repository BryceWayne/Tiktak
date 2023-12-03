package models

type Like struct {
    ID      string `json:"id"`
    VideoID string `json:"videoId"`
    UserID  string `json:"userId"`
}
