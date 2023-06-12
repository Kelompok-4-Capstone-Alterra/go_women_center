package handler

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction/usecase"
	"github.com/labstack/echo/v4"
)

type TransactionHandler interface {}

type transactionHandler struct {
	Usecase usecase.TransactionUsecase
}

func NewTransactionHandler(transactionUsecase usecase.TransactionUsecase) *transactionHandler {
	return &transactionHandler{}
}

func (h *transactionHandler) GenerateTransaction(c echo.Context) error {
	h.Usecase.GenerateTransaction()
	return nil
}