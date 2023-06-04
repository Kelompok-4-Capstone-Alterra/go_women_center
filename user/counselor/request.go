package counselor

type CreateReviewRequest struct {
	CounselorID string `param:"id" validate:"required,uuid"`
	UserID      string
	Rating      float32 `form:"rating" validate:"required,number,min=1,max=5"`
	Review      string  `form:"review" validate:"omitempty"`
}

type GetAllRequest struct {
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Search string `query:"search"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=high_rating low_price high_price"`
	Topic  int    `query:"topic" validate:"required,number,oneof=1 2 3 4 5 6 7 8 9 10"`
}

type GetAllReviewRequest struct {
	CounselorID string `param:"id" validate:"required,uuid"`
	Page        int    `query:"page"`
	Limit       int    `query:"limit"`
	Search      string `query:"search"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}