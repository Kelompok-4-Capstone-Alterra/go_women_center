package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction/usecase"
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

func (h *transactionHandler) GenerateTransaction() {

}

func (h *transactionHandler) SendTransaction(c echo.Context) error {
	// get jwt token and check for validity
	user := c.Get("user").(*helper.JwtCustomUserClaims)
	err := user.Valid()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	request := transaction.SendTransactionRequest{
		UserCredential: user,
	}
	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	data, err := h.Usecase.SendTransaction(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	return c.JSON(http.StatusOK, helper.ResponseData(
		"success creating new transaction",
		http.StatusOK,
		data,
	))
}

func (h *transactionHandler) GetAllTransaction(c echo.Context) error {
	// get jwt token and check for validity
	user := c.Get("user").(*helper.JwtCustomUserClaims)
	err := user.Valid()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	data, err := h.Usecase.GetAll(user.ID)
	
	return c.JSON(http.StatusOK, helper.ResponseData(
		"success get new transaction",
		http.StatusOK,
		data,
	))
}

func (h *transactionHandler) MidtransNotification(c echo.Context) error {
	notifMap := make(map[string]interface{})
	
	err := json.
		NewDecoder(c.Request().Body).
		Decode(&notifMap)
	if err != nil {
		return err
	}
	transactionStatus, ok := notifMap["transaction_status"].(string)
	if !ok {
		log.Println("error at trans_status")
		return c.JSON(http.StatusAccepted, "not a successfull transaction")
	}

	log.Println(transactionStatus)

	transactionId, ok := notifMap["order_id"].(string)
	if !ok {
		log.Println("error at trans_id")
		return c.JSON(http.StatusInternalServerError, transaction.ErrorTransactionNotFound)
	}

	log.Println(transactionId)

	// verify and update transaction status
	err = h.Usecase.UpdateStatus(transactionId, transactionStatus)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, transaction.ErrorTransactionNotFound)
	}

	return c.JSON(http.StatusOK, nil)
}
