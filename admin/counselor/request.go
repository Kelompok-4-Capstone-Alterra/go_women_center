package counselor

type CreateRequest struct {
	Name        string  `form:"name" validate:"required"`
	Email       string  `form:"email" validate:"required,email"`
	Username    string  `form:"username" validate:"required"`
	Topic       int     `form:"topic" validate:"required,number"`
	Description string  `form:"description" validate:"required"`
	Tarif       float64 `form:"tarif" validate:"required,number"`
}

type UpdateRequest struct {
	ID          string  `param:"id" validate:"required,uuid"`
	Name        string  `form:"name"`
	Email       string  `form:"email" validate:"omitempty,email"`
	Username    string  `form:"username"`
	Topic       int     `form:"topic" validate:"omitempty,number"`
	Description string  `form:"description"`
	Tarif       float64 `form:"tarif" validate:"omitempty,number"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}