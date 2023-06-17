package article

type CreateCommentRequest struct {
	ArticleID string `param:"id" validate:"required,uuid"`
	UserID    string
	Comment   string `form:"comment" validate:"required"`
}

type DeleteCommentRequest struct {
	ArticleID string `param:"article_id" validate:"required,uuid"`
	CommentID string `param:"comment_id" validate:"required,uuid"`
	UserID    string
}

type GetAllCommentRequest struct {
	ArticleID string `param:"id" validate:"required,uuid"`
}

type GetAllRequest struct {
	Search string `query:"search"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=newest oldest most_viewed"`
}

type UpdateCountRequest struct {
	ID           string `param:"id" validate:"required,uuid"`
	CommentCount string `validate:"omitempty"`
	ViewCount    string `validate:"omitempty"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}
