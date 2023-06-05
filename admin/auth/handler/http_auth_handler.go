package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	Usecase usecase.AuthUsecase
	JwtConf helper.AuthJWT
}

func NewAuthHandler(u usecase.AuthUsecase, jwtConf helper.AuthJWT) *authHandler {
	return &authHandler{
		Usecase: u,
		JwtConf: jwtConf,
	}
}

func (h *authHandler) LoginHandler(c echo.Context) error {
	request := auth.LoginAdminRequest{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	if err := isRequestValid(request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	data, err := h.Usecase.Login(request)
	if err != nil {
		status := http.StatusInternalServerError

		switch err {
			case auth.ErrInvalidCredential:
				status = http.StatusBadRequest
		}

		return c.JSON(status, helper.ResponseData(
			err.Error(),
			status,
			nil,
		))
	}

	token, _ := h.JwtConf.GenerateAdminToken(data.Email)

	return c.JSON(http.StatusOK, helper.ResponseData(
		"login success",
		http.StatusOK,
		echo.Map{
			"token": token,
		},
	))

}
