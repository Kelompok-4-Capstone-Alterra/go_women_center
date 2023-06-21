package career

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}

type SearchRequest struct {
	Search string `form:"search" validate:"omitempty"`
}

type GetAllRequest struct {
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Search string `query:"search"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=newest highest_salary"`
}