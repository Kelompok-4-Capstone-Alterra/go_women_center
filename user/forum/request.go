package forum

type GetAllRequest struct {
	UserId     string
	Topic      string `query:"topic" validate:"omitempty"`
	SortBy     string `query:"sort_by" validate:"omitempty,oneof=oldest newest popular"`
	CategoryId int    `query:"category_id" validate:"omitempty,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	MyForum    string `query:"my_forum" validate:"omitempty,oneof=yes"`
	Page       int    `query:"page" validate:"omitempty"`
	Offset     int    `query:"offset" validate:"omitempty"`
	Limit      int    `query:"limit" validate:"omitempty"`
}

type CreateRequest struct {
	ID         string
	UserId     string
	CategoryId int    `json:"category_id" form:"category_id" validate:"required,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	Link       string `json:"link" form:"link" validate:"required"`
	Topic      string `json:"topic" form:"topic" validate:"required"`
	Status     bool   `json:"status"`
	Member     int    `json:"member"`
}

type UpdateRequest struct {
	CategoryId int    `json:"category_id" form:"category_id" validate:"omitempty,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	Link       string `json:"link" form:"link" validate:"omitempty"`
	Topic      string `json:"topic" form:"topic" validate:"omitempty"`
}
