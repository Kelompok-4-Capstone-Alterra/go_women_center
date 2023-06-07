package profile

type IdRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}