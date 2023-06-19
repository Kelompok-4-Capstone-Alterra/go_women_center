package readingList

type GetAllRequest struct {
	UserId string
	Name   string `query:"name" validate:"omitempty"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=oldest newest"`
	Page   int    `query:"page" validate:"omitempty"`
	Offset int    `query:"offset" validate:"omitempty"`
	Limit  int    `query:"limit" validate:"omitempty"`
}

type CreateRequest struct {
	ID          string `gorm:"primarykey" json:"id" validate:"required"`
	UserId      string `json:"user_id" form:"user_id" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
}

type UpdateRequest struct {
	Name        string `json:"name" form:"name" validate:"omitempty"`
	Description string `json:"description" form:"description" validate:"omitempty"`
}
