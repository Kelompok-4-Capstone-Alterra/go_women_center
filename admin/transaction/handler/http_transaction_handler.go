package handler

import (
	// "net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)

type TransactionHandler interface{}

type transactionHandler struct {
	Usecase usecase.TransactionUsecase
}

func NewTransactionHandler(transactionUsecase usecase.TransactionUsecase) *transactionHandler {
	return &transactionHandler{
		Usecase: transactionUsecase,
	}
}

func (th *transactionHandler) GetAll(c echo.Context) error {
	code, data, err := th.Usecase.GetAll()
	if err != nil {
		return c.JSON(code, helper.ResponseData(
			err.Error(),
			code,
			nil,
		))
	}

	return c.JSON(code, helper.ResponseData(
		"success",
		code,
		data,
	))
}
