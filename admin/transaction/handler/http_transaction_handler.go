package handler

import (
	// "net/http"

	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction"
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
	getAllReq := transaction.GetAllRequest{}
	err := c.Bind(&getAllReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	page, offset, limit := helper.GetPaginateData(getAllReq.Page, getAllReq.Limit)

	code, totalPages, data, err := th.Usecase.GetAll(getAllReq.Search, getAllReq.SortBy, offset, limit)
	if err != nil {
		return c.JSON(code, helper.ResponseData(
			err.Error(),
			code,
			nil,
		))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get all transaction", http.StatusOK, echo.Map{
		"current_pages": page,
		"total_pages":   totalPages,
		"transaction":   data,
	}))
}

func (th *transactionHandler) SendLink(c echo.Context) error {
	req := transaction.SendLinkRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	// TODO: validate req

	code, err := th.Usecase.SendLink(req)
	if err != nil {
		return c.JSON(code, helper.ResponseData(
			err.Error(),
			code,
			nil,
		))
	}

	return c.JSON(code, helper.ResponseData(
		"success sending link",
		code,
		nil,
	))
}

func (th *transactionHandler) CancelTransaction(c echo.Context) error {
	req := transaction.CancelTransactionRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	// TODO: validate req

	code, err := th.Usecase.CancelTransaction(req)
	if err != nil {
		return c.JSON(code, helper.ResponseData(
			err.Error(),
			code,
			nil,
		))
	}

	return c.JSON(code, helper.ResponseData(
		"success canceling transaction",
		code,
		nil,
	))
}
