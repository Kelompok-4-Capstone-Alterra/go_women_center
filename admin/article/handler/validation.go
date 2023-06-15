package handler

import (
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/article"
	"github.com/go-playground/validator"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())

			if err.Tag() == "required" {
				return article.ErrRequired
			}

			switch field {
			case "email":
				return article.ErrEmailFormat
			case "id":
				return article.ErrIdFormat
			case "sort_by":
				return article.ErrInvalidSort
			}

		}
	}
	return nil
}
