package userForum

import "time"

type Response struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	ForumId   string    `json:"forum_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
