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
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	err = isRequestValid(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	err = isValidTopic(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	code, data, err := h.Usecase.SendTransaction(request)
	if err != nil {
		return c.JSON(code, helper.ResponseData(
			err.Error(),
			code,
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
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	return c.JSON(http.StatusOK, helper.ResponseData(
		"success get transaction",
		http.StatusOK,
		data,
	))
}

func (h *transactionHandler) GetTransactionDetail(c echo.Context) error {
	// get jwt token and check for validity
	user := c.Get("user").(*helper.JwtCustomUserClaims)
	err := user.Valid()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ResponseData(
			err.Error(),
			http.StatusUnauthorized,
			nil,
		))
	}

	transactionId := c.Param("id")

	code, data, err := h.Usecase.GetTransactionDetail(user.ID, transactionId)
	if err != nil {
		return c.JSON(code, helper.ResponseData(
			err.Error(),
			code,
			nil,
		))
	}

	return c.JSON(code, helper.ResponseData(
		"success get transaction",
		code,
		data,
	))
}

func (h *transactionHandler) UserJoinHandler(c echo.Context) error {
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
	
	req := transaction.UserJoinHandlerRequest{}
	err = c.Bind(&req)
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

	err = isValidUserId(req.UserId, user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ResponseData(
			err.Error(),
			http.StatusUnauthorized,
			nil,
		))
	}

	err = h.Usecase.UserJoinNotification(req.TransactionId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	return c.JSON(http.StatusOK, helper.ResponseData(
		"success update status",
		http.StatusOK,
		nil,
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
		return c.JSON(http.StatusOK, transaction.ErrorTransactionNotFound)
	}

	log.Println(transactionId)

	// verify and update transaction status
	err = h.Usecase.UpdateStatus(transactionId, transactionStatus)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusOK, transaction.ErrorTransactionNotFound)
	}

	return c.JSON(http.StatusOK, nil)
}
