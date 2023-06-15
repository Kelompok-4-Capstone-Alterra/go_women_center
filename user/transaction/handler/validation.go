package handler

import (
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
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
			case "counselor_id":
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