package handler

import (
	"net/http"
	"strconv"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/users"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/users/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetById(c echo.Context) error {
	
	var req users.IdRequest
	c.Bind(&req)

	if err := isRequestValid(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(),http.StatusBadRequest, nil))
	}
	
	userRes, err := h.usecase.GetById(req.ID)
	if err != nil {
		status := http.StatusInternalServerError

		if err == users.ErrUserNotFound {
			status = http.StatusNotFound
		}

		return c.JSON(status, helper.ResponseData(err.Error(),status, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get user by id",http.StatusOK, echo.Map{
		"user": userRes,
	}))
}

func (h *UserHandler) GetAll(c echo.Context) error {
	page, _ :=  strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	search := c.QueryParam("search")

	page, offset, limit := helper.GetPaginateData(page, limit)

	usersRes, totalPages, err := h.usecase.GetAll(search, offset, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(err.Error(),http.StatusInternalServerError, nil))
	}

	if page > totalPages {
		return c.JSON(http.StatusNotFound, helper.ResponseData(users.ErrPageNotFound.Error(),http.StatusNotFound, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get all user",http.StatusOK, echo.Map{
		"users": usersRes,
		"current_pages": page,
		"total_pages": totalPages,
	}))
}

func (h *UserHandler) Delete(c echo.Context) error {
	var req users.IdRequest
	c.Bind(&req)

	if err := isRequestValid(req); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(),http.StatusBadRequest, nil))
	}

	err := h.usecase.Delete(req.ID)
	if err != nil {
		status := http.StatusInternalServerError

		if err == users.ErrUserNotFound {
			status = http.StatusNotFound
		}

		return c.JSON(status, helper.ResponseData(err.Error(),status, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success delete user",http.StatusOK, nil))
}
