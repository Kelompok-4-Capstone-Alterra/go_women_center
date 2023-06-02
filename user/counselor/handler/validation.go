package handler

import (
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor"
	"github.com/go-playground/validator"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())

			if err.Tag() == "required" {
				return counselor.ErrRequired
			}

			if err.Tag() == "uuid" {
				return counselor.ErrIdFormat
			}

			switch field {
			case "tarif":
				return counselor.ErrTarifFormat
			case "rating":
				return counselor.ErrRatingFormat
			}
			

		}
	}

	return nil
}