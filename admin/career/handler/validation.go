package handler

import (
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/career"
	"github.com/go-playground/validator"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())

			if err.Tag() == "required" {
				return career.ErrRequired
			}

			switch field {
			case "email":
				return career.ErrEmailFormat
			case "id":
				return career.ErrIdFormat
			}

		}
	}
	return nil
}
