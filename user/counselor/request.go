package counselor

type CreateReviewRequest struct {
	TransactionID string `json:"transaction_id" validate:"required,uuid"`
	CounselorID   string `param:"id" validate:"required,uuid"`
	UserID        string
	Rating        float32 `json:"rating" validate:"required,number,min=1,max=5"`
	Review        string  `json:"review" validate:"omitempty"`
}

type GetAllRequest struct {
	Search      string `query:"search"`
	SortBy      string `query:"sort_by" validate:"omitempty,oneof=highest_rating lowest_price highest_price"`
	Topic       int    `query:"topic" validate:"omitempty,number,oneof=1 2 3 4 5 6 7 8 9 10"`
	IsAvailable string `query:"is_available" validate:"omitempty,oneof=true"`
}

type GetAllReviewRequest struct {
	CounselorID string `param:"id" validate:"required,uuid"`
	Page        int    `query:"page"`
	Limit       int    `query:"limit"`
	Search      string `query:"search"`
}

type IdRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}