package handler

import (
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction"
	"github.com/go-playground/validator"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())

			if err.Tag() == "required" {
				return transaction.ErrRequired
			}

			switch field {
			case "counselorid":
				return transaction.ErrInvalidUUID
			case "status":
				return transaction.ErrorInvalidPaymentStatus
			case "transactionid":
				return transaction.ErrInvalidUUID
			}
		}
	}

	return nil
}

func isValidTopic(req transaction.SendTransactionRequest) error {
	if !constant.METHOD[req.ConsultationMethod] {
		return transaction.ErrInvalidConsultationMethod
	}
	return nil
}

func isValidUserId(user_id string, token *helper.JwtCustomUserClaims) error {
	if user_id != token.ID {
		return transaction.ErrInvalidUserCredential
	}
	return nil
}
