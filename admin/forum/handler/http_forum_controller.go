package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo"
)

type ForumHandler struct {
	ForumU entity.ForumUsecase
}

func NewForumHandler(ForumU entity.ForumUsecase) entity.ForumHandler {
	return &ForumHandler{
		ForumU: ForumU,
	}
}

func (fh ForumHandler) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", entity.Forum{}))
}

func (fh ForumHandler) GetById(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", entity.Forum{}))
}

func (fh ForumHandler) Create(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", entity.Forum{}))
}

func (fh ForumHandler) Update(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", entity.Forum{}))
}

func (fh ForumHandler) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", entity.Forum{}))
}
