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
	
	var user = c.Get("user").(*helper.JwtCustomUserClaims)

	profile, err := h.usecase.GetById(user.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(err.Error(), http.StatusInternalServerError, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get profile", http.StatusOK, echo.Map{
		"profile": profile,
	}))
}


func(h *ProfileHandler) Update(c echo.Context) error {
	
	var user = c.Get("user").(*helper.JwtCustomUserClaims)

	var req profile.UpdateRequest
	file, _ := c.FormFile("profile_picture")
	req.ProfilePicture = file

	c.Bind(&req)

	helper.RemoveWhiteSpace(&req)

	if err := isRequestValid(req) ;err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(err.Error(), http.StatusInternalServerError, nil))
	}

	req.ID = user.ID
	err := h.usecase.Update(req)

	if err != nil {
		status := http.StatusInternalServerError

		switch err {
			case profile.ErrUserNotFound:
				status = http.StatusNotFound
			case profile.ErrUsernameDuplicate:
				status = http.StatusConflict
			case
				profile.ErrProfilePictureFormat:
				status = http.StatusBadRequest
		}

		return c.JSON(status, helper.ResponseData(err.Error(), status, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success update profile", http.StatusOK, nil))
}

func(h *ProfileHandler) UpdatePassword(c echo.Context) error {
	
	var user = c.Get("user").(*helper.JwtCustomUserClaims)

	var req profile.UpdatePasswordRequest

	c.Bind(&req)

	helper.RemoveWhiteSpace(&req)

	if err := isRequestValid(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	req.ID = user.ID
	err := h.usecase.UpdatePassword(req)

	if err != nil {
		
		status := http.StatusInternalServerError

		switch err {
			case profile.ErrPasswordNotMatch:
				status = http.StatusBadRequest
		}

		return c.JSON(status, helper.ResponseData(err.Error(), status, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success update password", http.StatusOK, nil))
}