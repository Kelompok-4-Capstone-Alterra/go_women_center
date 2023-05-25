package counselor

type CreateRequest struct {
	FullName       string  `form:"full_name" validate:"required"`
	Email          string  `form:"email" validate:"required,email"`
	Username       string  `form:"username" validate:"required"`
	Topic          string  `form:"topic" validate:"required"`
	Description    string  `form:"description" validate:"required"`
	Tarif          float64 `form:"tarif" validate:"required,number"`
	ProfilePicture string
}
