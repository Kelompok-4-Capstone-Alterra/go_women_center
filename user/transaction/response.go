package transaction

type SendTransactionResponse struct {
	TransactionID string `json:"transaction_id"`
	PaymentLink   string `json:"payment_link"`
	Data          interface{}
}
