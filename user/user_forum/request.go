package userForum

type CreateRequest struct {
	ID      string
	UserId  string
	ForumId string `json:"forum_id" validate:"required"`
}
