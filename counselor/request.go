package counselor

type CreateRequest struct {
	FullName    string  `form:"full_name" validate:"required"`
	Email       string  `form:"email" validate:"required,email"`
	Username    string  `form:"username" validate:"required"`
	Topic       string  `form:"topic" validate:"required"`
	Description string  `form:"description" validate:"required"`
	Tarif       float64 `form:"tarif" validate:"required,number"`
}

type UpdateRequest struct {
	ID          string  `param:"id" validate:"required,uuid"`
	FullName    string  `form:"full_name"`
	Email       string  `form:"email" validate:"omitempty,email"`
	Username    string  `form:"username"`
	Topic       string  `form:"topic"`
	Description string  `form:"description"`
	Tarif       float64 `form:"tarif" validate:"omitempty,number"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}