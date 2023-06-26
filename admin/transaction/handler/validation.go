package handler

import (
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction"
	"github.com/go-playground/validator"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())

			if err.Tag() == "required" {
				return transaction.ErrRequired
			}

			switch field {
			case "page":
				return transaction.ErrPage
			case "limit":
				return transaction.ErrLimit
			case "search":
				return transaction.ErrSearch
			case "sortby":
				return transaction.ErrSortBy
			case "transactionid":
				return transaction.ErrTransactionId
			case "link":
				return transaction.ErrInvalidLink
			case "startdate":
				return transaction.ErrInvalidStartDate
			case "enddate":
				return transaction.ErrInvalidEndDate
			}

		}
	}

	return nil
}

var validUrlHost = map[string]bool{
	"us05web.zoom.us": true,
	"zoom.us": true,
	"t.me": true,
}