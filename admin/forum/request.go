package forum

type GetAllRequest struct {
	Topic      string `query:"topic" validate:"omitempty"`
	SortBy     string `query:"sort_by" validate:"omitempty,oneof=oldest newest popular"`
	CategoryId int    `query:"category_id" validate:"omitempty,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	Page       int    `query:"page" validate:"omitempty"`
	Offset     int    `query:"offset" validate:"omitempty"`
	Limit      int    `query:"limit" validate:"omitempty"`
}
