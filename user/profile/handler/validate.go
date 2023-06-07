package handler

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/profile"
	"github.com/go-playground/validator"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// field := strings.ToLower(err.Field())
			
			if err.Tag() == "required" {
				return profile.ErrRequired
			}

			if err.Tag() == "uuid" {
				return profile.ErrIdFormat
			}


		}
	}

	return nil
}