package middleware

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/usecase"
	"github.com/labstack/echo/v4"
)

func CheckUser(usecase usecase.UserUsecase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var user = c.Get("user").(*helper.JwtCustomUserClaims)

			_, err :=  usecase.GetById(user.ID)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, helper.ResponseData(auth.ErrInvalidCredential.Error(), http.StatusUnauthorized, nil))
			}

			return next(c)
		}
	}
}