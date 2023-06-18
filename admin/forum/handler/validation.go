package handler

import (
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum"
	"github.com/go-playground/validator"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())

			if err.Tag() == "required" {

				if field == "category" {
					return forum.ErrRequiredCategory
				} else if field == "link" {
					return forum.ErrRequiredLink
				} else if field == "topic" {
					return forum.ErrRequiredTopic
				}

				return forum.ErrRequired
			}

			switch field {
			case "sortby":
				return forum.ErrInvalidSort
			case "category":
				return forum.ErrInvalidCategory
			}

		}
	}

	return nil
}
