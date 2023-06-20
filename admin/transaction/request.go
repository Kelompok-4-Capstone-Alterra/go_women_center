package transaction

type GetAllRequest struct {
	Page   int    `query:"page" validate:"omitempty"`
	Limit  int    `query:"limit" validate:"omitempty"`
	Search string `query:"search"`
	SortBy string `query:"sort_by"`
}

type SendLinkRequest struct {
	TransactionId string `json:"transaction_id" validate:"required,uuid"`
	Link          string `json:"link" validate:"required"`
}

type CancelTransactionRequest struct {
	TransactionId string `json:"transaction_id" validate:"required,uuid"`
}

type ReportRequest struct {
	StartDate  string    `query:"start_date" validate:"omitempty"`
	EndDate    string    `query:"end_date" validate:"omitempty"`
	Page       int    `query:"page" validate:"omitempty"`
	Limit      int    `query:"limit" validate:"omitempty"`
	Search     string `query:"search"`
	SortBy     string `query:"sort_by"`
	Offset     int
	IsDownload bool
}
