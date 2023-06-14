package article

type GetAllResponse struct {
	ID           string `json:"id"`
	Image        string `json:"image"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	Topic        string `json:"topic"`
	ViewCount    int    `json:"view_count"`
	CommentCount int    `json:"comment_count"`
	Date         string `json:"date"`
	Saved        bool   `json:"saved"`
}

type GetByResponse struct {
	ID           string `json:"id"`
	Image        string `json:"image"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	Topic        string `json:"topic"`
	ViewCount    int    `json:"view_count"`
	CommentCount int    `json:"comment_count"`
	Description  string `json:"description"`
	Date         string `json:"date"`
}

type CommentResponse struct {
	ID             string `json:"id"`
	ArticleID      string `json:"article_id"`
	UserID         string `json:"user_id"`
	ProfilePicture string `json:"profile_picture"`
	Username       string `json:"username"`
	Comment        string `json:"comment"`
	CreatedAt      string `json:"created_at"`
}
