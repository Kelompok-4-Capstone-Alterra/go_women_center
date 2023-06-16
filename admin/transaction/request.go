package transaction

type SendLinkRequest struct {
	TransactionId string `json:"transaction_id" validate:"required,uuid"`
	Link          string `json:"link" validate:"required"`
}

type CancelTransactionRequest struct {
	TransactionId string `json:"transaction_id" validate:"required,uuid"`
}
