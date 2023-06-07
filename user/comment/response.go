package comment

import "time"

type GetAllResponse struct {
	ID        string    `json:"id"`
	ArticleID string    `json:"article_id"`
	UserID    string    `json:"user_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

type GetByCounselorId struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
