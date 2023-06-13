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
		"success get all conselor",
		http.StatusOK,
		echo.Map{
			"data": data,
		},
	))
}

func (h *transactionHandler) Notification(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.
		NewDecoder(c.Request().Body).
		Decode(&json_map)
	if err != nil {
		return err
	}

	log.Println(json_map)

	return c.JSON(http.StatusOK, nil)
}
