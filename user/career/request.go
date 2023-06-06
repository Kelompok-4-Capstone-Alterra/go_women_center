package career

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}

type SearchRequest struct {
	Search string `form:"search" validate:"omitempty"`
}