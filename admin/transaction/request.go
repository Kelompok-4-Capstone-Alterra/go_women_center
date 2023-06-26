package transaction

type GetAllRequest struct {
	Page   int    `query:"page" validate:"omitempty"`
	Limit  int    `query:"limit" validate:"omitempty"`
	Search string `query:"search" validate:"omitempty"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=oldest newest"`
}

type SendLinkRequest struct {
	TransactionId string `json:"transaction_id" validate:"required,uuid"`
	Link          string `json:"link" validate:"required"`
}

type CancelTransactionRequest struct {
	TransactionId string `json:"transaction_id" validate:"required,uuid"`
}

// after initialization, set the download
type ReportRequest struct {
	StartDate  string `query:"start_date" validate:"omitempty"`
	EndDate    string `query:"end_date" validate:"omitempty"`
	Page       int    `query:"page" validate:"omitempty"`
	Limit      int    `query:"limit" validate:"omitempty"`
	Search     string `query:"search" validate:"omitempty"`
	SortBy     string `query:"sort_by" validate:"omitempty,oneof=oldest newest"`
	Offset     int
	IsDownload bool
}
