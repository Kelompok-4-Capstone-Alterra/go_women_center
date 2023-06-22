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

	err = isRequestValid(getAllReq)
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

	err = isRequestValid(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	err = helper.IsValidUrl(req.Link, validUrlHost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

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

	err = isRequestValid(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

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

func (th *transactionHandler) GetReport(c echo.Context) error {
	reportReq := transaction.ReportRequest{}
	err := c.Bind(&reportReq)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}
	reportReq.IsDownload = false

	err = isRequestValid(reportReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	page, offset, limit := helper.GetPaginateData(reportReq.Page, reportReq.Limit)
	reportReq.Page = page
	reportReq.Offset = offset
	reportReq.Limit = limit

	data, totalPages, err := th.Usecase.GetAllForReport(reportReq)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get all transaction", http.StatusOK, echo.Map{
		"current_pages": page,
		"total_pages":   totalPages,
		"transaction":   data,
	}))
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

	err = isRequestValid(reportReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	log.Println(reportReq)

	data, _, err := th.Usecase.GetAllForReport(reportReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	csvData, status, err := th.Usecase.GenerateReport(data)
	if err != nil {
		return c.JSON(status, helper.ResponseData(
			err.Error(),
			status,
			nil,
		))
	}

	// Set the appropriate headers for the CSV response
    c.Response().Header().Set("Content-Type", "text/csv")
    c.Response().Header().Set("Content-Disposition", "attachment; filename=export.csv")

    // Write the CSV data as the response body
    return c.Blob(status, "text/csv", csvData)

}
