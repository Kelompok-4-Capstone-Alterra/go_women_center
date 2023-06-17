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
