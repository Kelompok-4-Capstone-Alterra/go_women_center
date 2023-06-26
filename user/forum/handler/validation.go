package handler

import (
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
	"github.com/go-playground/validator"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())

			if err.Tag() == "required" {

				if field == "categoryid" {
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
			case "categoryid":
				return forum.ErrInvalidCategory
			case "myforum":
				return forum.ErrInvalidMyForum
			}

		}
	}

	return nil
}

var validUrlHost = map[string]bool{
	"t.me": true,
}
