package article

type CreateCommentRequest struct {
	ArticleID string `param:"id" validate:"required,uuid"`
	UserID    string
	Comment   string `form:"comment" validate:"required"`
}

type DeleteCommentRequest struct {
	ArticleID string `param:"id" validate:"required,uuid"`
	UserID    string
}

type GetAllCommentRequest struct {
	ArticleID string `param:"id" validate:"required,uuid"`
	Page      int    `query:"page"`
	Limit     int    `query:"limit"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}
