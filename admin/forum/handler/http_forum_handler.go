package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)

type ForumHandlerInterface interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Delete(c echo.Context) error
}

type ForumHandler struct {
	ForumU usecase.ForumUsecaseInterface
}

func NewForumHandler(ForumU usecase.ForumUsecaseInterface) ForumHandlerInterface {
	return &ForumHandler{
		ForumU: ForumU,
	}
}

func (fh ForumHandler) GetAll(c echo.Context) error {
	var getAllParam forum.GetAllRequest
	c.Bind(&getAllParam)
	getAllParam.Page, getAllParam.Offset, getAllParam.Limit = helper.GetPaginateData(getAllParam.Page, getAllParam.Limit)

	forums, totalPages, err := fh.ForumU.GetAll(getAllParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	if getAllParam.Page > totalPages {
		return c.JSON(http.StatusBadRequest, helper.ResponseData("page not found", http.StatusBadRequest, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success to get all forum data", http.StatusOK, echo.Map{
		"forums":        forums,
		"current_pages": getAllParam.Page,
		"total_pages":   totalPages,
	}))
}

func (fh ForumHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	forum, err := fh.ForumU.GetById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("success to get forum data details", http.StatusOK, forum))
}

func (fh ForumHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := fh.ForumU.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("successfully deleted forum data", http.StatusOK, nil))
}
