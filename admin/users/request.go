package users

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}

type GetAllRequest struct {
	Page   int    `query:"page" validate:"omitempty"`
	Limit  int    `query:"limit" validate:"omitempty"`
	Search string `query:"search" validate:"omitempty"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=oldest newest"`
}