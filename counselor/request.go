package counselor

type CreateRequest struct {
	Name        string  `form:"name" validate:"required"`
	Email       string  `form:"email" validate:"required,email"`
	Username    string  `form:"username" validate:"required"`
	Topic       string  `form:"topic" validate:"required"`
	Description string  `form:"description" validate:"required"`
	Tarif       float64 `form:"tarif" validate:"required,number"`
}

type UpdateRequest struct {
	ID          string  `param:"id" validate:"required,uuid"`
	Name        string  `form:"name"`
	Email       string  `form:"email" validate:"omitempty,email"`
	Username    string  `form:"username"`
	Topic       string  `form:"topic"`
	Description string  `form:"description"`
	Tarif       float64 `form:"tarif" validate:"omitempty,number"`
}

type CreateReviewRequest struct {
	CounselorID string  `param:"id" validate:"required,uuid"`
	UserID      string  `form:"user_id" validate:"required,uuid"`
	Comment     string  `form:"comment" validate:"omitempty"`
	Rating      float32 `form:"rating" validate:"required,number,min=1,max=5"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}