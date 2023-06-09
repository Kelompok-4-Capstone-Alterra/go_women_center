package handler

import (
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor"
	"github.com/go-playground/validator"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())

			if err.Tag() == "required" {

				if field == "topic" {
					return counselor.ErrRequiredTopic
				}
				
				return counselor.ErrRequired
			}

			switch field {
			case "hasschedule": 
				return counselor.ErrHasScheduleFormat
			case "email":
				return counselor.ErrEmailFormat
			case "price":
				return counselor.ErrPriceFormat
			case "rating":
				return counselor.ErrRatingFormat
			case "id":
				return counselor.ErrIdFormat
			case "topic":
				return counselor.ErrInvalidTopic
			case "sortby":
				return counselor.ErrInvalidSort
			}
			

		}
	}

	return nil
}
