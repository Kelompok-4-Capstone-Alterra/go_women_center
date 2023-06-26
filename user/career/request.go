package career

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}

type SearchRequest struct {
	Search string `form:"search" validate:"omitempty"`
}

type GetAllRequest struct {
	Search string `query:"search"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=newest highest_salary"`
}