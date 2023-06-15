package readingList

type GetAllRequest struct {
	UserId string
	Name   string `query:"name"`
	Sort   string `query:"sort"`
	Page   int    `query:"page"`
	Offset int    `query:"offset"`
	Limit  int    `query:"limit"`
}

type CreateRequest struct {
	ID          string `gorm:"primarykey" json:"id"`
	UserId      string `json:"user_id" form:"user_id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

type UpdateRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}
