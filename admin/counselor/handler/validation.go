package handler

import (
	"mime/multipart"
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
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
				return counselor.ErrRequired
			}

			switch field {
				case "email":
					return counselor.ErrEmailFormat
				case "tarif":
					return counselor.ErrTarifFormat
				case "rating":
					return counselor.ErrRatingFormat
				case "id":
					return counselor.ErrIdFormat
			}

		}
	}

	if data, ok := m.(counselor.CreateRequest); ok {
		
		if _, ok := constant.TOPICS[data.Topic]; !ok {
			return counselor.ErrInvalidTopic
		}
	}

	if data, ok := m.(counselor.UpdateRequest); ok {
		if _, ok := constant.TOPICS[data.Topic]; !ok {
			return counselor.ErrInvalidTopic
		}
	}

	return nil
}

func isImageValid(img *multipart.FileHeader) error {

	if img == nil {
		return counselor.ErrRequired
	}
	
	if img.Size > 2 * 1024 * 1024 { // 2 MB
		return counselor.ErrProfilePictureSize
	}

	if !helper.IsImageValid(img) {
		return counselor.ErrProfilePictureFormat
	}

	return nil
}	