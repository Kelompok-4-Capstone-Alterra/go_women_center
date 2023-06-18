package forum

type GetAllRequest struct {
	Topic    string `query:"topic" validate:"omitempty"`
	SortBy   string `query:"sort_by" validate:"omitempty,oneof=oldest newest popular"`
	Category int    `query:"category" validate:"omitempty,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	Page     int    `query:"page" validate:"omitempty"`
	Offset   int    `query:"offset" validate:"omitempty"`
	Limit    int    `query:"limit" validate:"omitempty"`
}

type CreateRequest struct {
	ID       string
	UserId   string
	Category int    `json:"category" form:"category" validate:"required,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	Link     string `json:"link" form:"link" validate:"required"`
	Topic    string `json:"topic" form:"topic" validate:"required"`
	Status   bool   `json:"status"`
	Member   int    `json:"member"`
}

type UpdateRequest struct {
	Category int    `json:"category" form:"category" validate:"omitempty"`
	Link     string `json:"link" form:"link" validate:"omitempty"`
	Topic    string `json:"topic" form:"topic" validate:"omitempty"`
}
