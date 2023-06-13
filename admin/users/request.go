package users

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}
