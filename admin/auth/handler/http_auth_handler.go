package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	Usecase usecase.AuthUsecase
	JwtConf helper.AuthJWT
}

func NewUserHandler(u usecase.AuthUsecase, jwtConf helper.AuthJWT) *authHandler {
	return &authHandler{
		Usecase: u,
		JwtConf: jwtConf,
	}
}

func (h *authHandler) Login(c echo.Context) error {
	request := auth.LoginAdminDTO{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "login failed",
			"error": err.Error(),
		})	
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "login successfull",
	})
}
