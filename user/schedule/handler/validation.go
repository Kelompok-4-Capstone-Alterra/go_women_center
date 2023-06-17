package handler

import (
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/schedule"
	"github.com/go-playground/validator"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			
			if err.Tag() == "required" {
				switch field {
				case "counselorid":
					return schedule.ErrIdRequired
				}
			}
			switch field {
			case "counselorid":
				return schedule.ErrIdFormat
			}

		}
	}

	return nil
}
