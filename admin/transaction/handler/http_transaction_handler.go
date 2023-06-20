package handler

import (
	"log"
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

	status, totalPages, data, err := th.Usecase.GetAll(getAllReq.Search, getAllReq.SortBy, offset, limit)
	if err != nil {
		return c.JSON(status, helper.ResponseData(
			err.Error(),
			status,
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

	status, err := th.Usecase.SendLink(req)
	if err != nil {
		return c.JSON(status, helper.ResponseData(
			err.Error(),
			status,
			nil,
		))
	}

	return c.JSON(status, helper.ResponseData(
		"success sending link",
		status,
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

	status, err := th.Usecase.CancelTransaction(req)
	if err != nil {
		return c.JSON(status, helper.ResponseData(
			err.Error(),
			status,
			nil,
		))
	}

	return c.JSON(status, helper.ResponseData(
		"success canceling transaction",
		status,
		nil,
	))
}

func (th *transactionHandler) DownloadReport(c echo.Context) error {
	// TODO: validation
	reportReq := transaction.ReportRequest{}
	err := c.Bind(&reportReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}
	reportReq.IsDownload = true

	log.Println(reportReq)

	data, _, err := th.Usecase.GetAllForReport(reportReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	fileLocation, status, err := th.Usecase.GenerateReport(data)
	if err != nil {
		return c.JSON(status, helper.ResponseData(
			err.Error(),
			status,
			nil,
		))
	}

	return c.Attachment(fileLocation, "report.csv")
}
