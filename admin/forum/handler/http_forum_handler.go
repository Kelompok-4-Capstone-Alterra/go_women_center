package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)

type ForumHandlerInterface interface {
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

func (fh ForumHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := fh.ForumU.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData("failed to delete forum data", http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("successfully deleted forum data", http.StatusOK, nil))
}
