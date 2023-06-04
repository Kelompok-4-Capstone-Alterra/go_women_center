package handler

import (
	"net/http"
	"strings"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ForumHandlerInterface interface {
	GetAll(c echo.Context) error
	GetByCategory(c echo.Context) error
	GetByMyForum(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
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
	forums, err := fh.ForumU.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(http.StatusBadRequest, "Failed to get all forums data", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "Success to get all forums data", forums))
}

func (fh ForumHandler) GetByCategory(c echo.Context) error {
	id_category := c.Param("id")
	forums, err := fh.ForumU.GetByCategory(id_category)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(http.StatusBadRequest, "Failed to get all forums data by category", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "Success to get all forums data by category", forums))
}

func (fh ForumHandler) GetByMyForum(c echo.Context) error {
	id_user := "1"
	forums, err := fh.ForumU.GetByMyForum(id_user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(http.StatusBadRequest, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "Success to get all forums data by user", forums))
}

func (fh ForumHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	forum, err := fh.ForumU.GetById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(http.StatusBadRequest, "Failed to get detail forums data", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "Success to get detail forum data", forum))
}

func (fh ForumHandler) Create(c echo.Context) error {
	var forum entity.Forum
	c.Bind(&forum)

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	forum.ID = uuid

	err := fh.ForumU.Create(&forum)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(http.StatusBadRequest, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "Successfully created forum data", nil))
}

func (fh ForumHandler) Update(c echo.Context) error {
	forum := entity.Forum{}
	id := c.Param("id")
	c.Bind(&forum)
	err := fh.ForumU.Update(id, &forum)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(http.StatusBadRequest, "Failed update forums", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "Successfully updated forum data", nil))
}

func (fh ForumHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := fh.ForumU.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(http.StatusBadRequest, "Failed get all forums", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "Successfully deleted forum data", nil))
}
