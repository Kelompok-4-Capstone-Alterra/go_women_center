package handler

import (
	"strings"

	userForum "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum"
	"github.com/go-playground/validator"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())

			if err.Tag() == "required" {

				if field == "forumid" {
					return userForum.ErrRequiredForumId
				}
				return userForum.ErrRequired
			}

		}
	}

	return nil
}
