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
	request := auth.LoginAdminDTO{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			http.StatusInternalServerError,
			err.Error(),
			nil,
		))
	}

	data, err := h.Usecase.Login(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			http.StatusInternalServerError,
			err.Error(),
			nil,
		))
	}

	token, err := h.JwtConf.GenerateAdminToken(data.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			http.StatusInternalServerError,
			err.Error(),
			nil,
		))
	}

	return c.JSON(http.StatusOK, helper.ResponseData(
		http.StatusOK,
		"login success",
		map[string]interface{}{
			"token": token,
		},
	))
}
