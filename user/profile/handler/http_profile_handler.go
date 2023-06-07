package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/profile"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/profile/usecase"
	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	usecase usecase.ProfileUsecase
}

func NewProfileHandler(PUscase usecase.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{PUscase}
}

func(h *ProfileHandler) GetById(c echo.Context) error {
	
	var req profile.IdRequest

	c.Bind(&req)

	if err := isRequestValid(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	profile, err := h.usecase.GetById(req.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(err.Error(), http.StatusInternalServerError, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get profile", http.StatusOK, profile))
}