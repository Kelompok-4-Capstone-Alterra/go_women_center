package forum

type GetAllRequest struct {
	UserId   string
	Topic    string `query:"topic" validate:"omitempty"`
	SortBy   string `query:"sort_by" validate:"omitempty,oneof=oldest newest popular"`
	Category int    `query:"category" validate:"omitempty,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	MyForum  string `query:"my_forum" validate:"omitempty,oneof=oldest newest popular"`
	Page     int    `query:"page" validate:"omitempty"`
	Offset   int    `query:"offset" validate:"omitempty"`
	Limit    int    `query:"limit" validate:"omitempty"`
}

type CreateRequest struct {
	ID       string
	UserId   string
	Category int    `json:"category" form:"category"`
	Link     string `json:"link" form:"link"`
	Topic    string `json:"topic" form:"topic"`
	Status   bool   `json:"status"`
	Member   int    `json:"member"`
}

type UpdateRequest struct {
	Category int    `json:"category" form:"category"`
	Link     string `json:"link" form:"link"`
	Topic    string `json:"topic" form:"topic"`
}
