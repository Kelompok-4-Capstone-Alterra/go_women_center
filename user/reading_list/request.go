package readingList

type GetAllRequest struct {
	IdUser     string
	Topic      string `query:"topic"`
	Created    string `query:"created"`
	Popular    string `query:"popular"`
	Categories int    `query:"categories"`
	MyForum    string `query:"myforum"`
	Page       int    `query:"page"`
	Offset     int    `query:"offset"`
	Limit      int    `query:"limit"`
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
