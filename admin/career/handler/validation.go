package handler

import (
	"mime/multipart"
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/career"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
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

func isImageValid(img *multipart.FileHeader) error {

	if img == nil {
		return career.ErrRequired
	}
	
	if img.Size > 10 * 1024 * 1024 { // 10 MB
		return career.ErrImageFormat
	}

	if !helper.IsImageValid(img) {
		return career.ErrImageSize
	}

	return nil
}