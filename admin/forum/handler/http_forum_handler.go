package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)

type ForumAdminHandlerInterface interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Delete(c echo.Context) error
}

type ForumAdminHandler struct {
	ForumAdminU usecase.ForumAdminUsecaseInterface
}

func NewForumAdminHandler(ForumAdminU usecase.ForumAdminUsecaseInterface) ForumAdminHandlerInterface {
	return &ForumAdminHandler{
		ForumAdminU: ForumAdminU,
	}
}

func (fah ForumAdminHandler) GetAll(c echo.Context) error {
	var getAllRequest forum.GetAllRequest
	c.Bind(&getAllRequest)
	getAllRequest.Page, getAllRequest.Offset, getAllRequest.Limit = helper.GetPaginateData(getAllRequest.Page, getAllRequest.Limit)

	if err := isRequestValid(getAllRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	forums, totalPages, err := fah.ForumAdminU.GetAll(getAllRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	if getAllRequest.Page > totalPages {
		return c.JSON(http.StatusNotFound, helper.ResponseData("page not found", http.StatusNotFound, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("success to get all forum data", http.StatusOK, echo.Map{
		"forums":        forums,
		"current_pages": getAllRequest.Page,
		"total_pages":   totalPages,
	}))
}

func (fah ForumAdminHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	forum, err := fah.ForumAdminU.GetById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("success to get forum data details", http.StatusOK, forum))
}

func (fah ForumAdminHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := fah.ForumAdminU.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("successfully deleted forum data", http.StatusOK, nil))
}
