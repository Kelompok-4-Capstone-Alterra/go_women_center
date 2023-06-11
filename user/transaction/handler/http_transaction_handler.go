package handler

type TransactionHandler interface {}

type transactionHandler struct {}

func NewTransactionHandler() *transactionHandler {
	return &transactionHandler{}
}