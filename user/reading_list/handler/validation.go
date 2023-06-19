package handler

import (
	"strings"

	readingList "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list"
	"github.com/go-playground/validator"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())

			if err.Tag() == "required" {

				if field == "id" {
					return readingList.ErrRequiredId
				} else if field == "userid" {
					return readingList.ErrRequiredUserId
				} else if field == "name" {
					return readingList.ErrRequiredName
				} else if field == "description" {
					return readingList.ErrRequiredDescription
				}

				return readingList.ErrRequired
			}

			switch field {
			case "sortby":
				return readingList.ErrInvalidSort
			}

		}
	}

	return nil
}
