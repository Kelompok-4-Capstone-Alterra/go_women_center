package counselor

type CreateReviewRequest struct {
	CounselorID string `param:"id" validate:"required,uuid"`
	UserID      string
	Rating      float32 `form:"rating" validate:"required,number,min=1,max=5"`
	Review      string  `form:"review" validate:"omitempty"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}