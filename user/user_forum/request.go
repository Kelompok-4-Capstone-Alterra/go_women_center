package userForum

type CreateRequest struct {
	ID      string `json:"id" validate:"required"`
	UserId  string `json:"user_id" validate:"required"`
	ForumId string `json:"forum_id" validate:"required"`
}
