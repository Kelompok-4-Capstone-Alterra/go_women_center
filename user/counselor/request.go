package counselor

type CreateReviewRequest struct {
	CounselorID string  `param:"id" validate:"required,uuid"`
	UserID      string  `form:"user_id" validate:"required,uuid"`
	Comment     string  `form:"comment" validate:"omitempty"`
	Rating      float32 `form:"rating" validate:"required,number,min=1,max=5"`
}