package article

import "time"

type GetAllResponse struct {
	ID string `json:"id"`
	Image        string    `json:"image"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	Topic        string    `json:"topic"`
	ViewCount    int       `json:"view_count"`
	CommentCount int       `json:"comment_count"`
}

type GetByResponse struct {
	ID           string    `json:"id"`
	Image        string    `json:"image"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	Topic        string    `json:"topic"`
	ViewCount    int       `json:"view_count"`
	CommentCount int       `json:"comment_count"`
	Description  string    `json:"description"`
	Date         time.Time `json:"date"`
}
