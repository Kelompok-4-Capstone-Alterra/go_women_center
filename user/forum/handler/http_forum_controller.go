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
		return c.JSON(http.StatusBadRequest, helper.ResponseData(http.StatusBadRequest, "Failed get all forums", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", forums))
}

func (fh ForumHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	forum, err := fh.ForumU.GetById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(http.StatusBadRequest, "Failed get all forums", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", forum))
}

func (fh ForumHandler) Create(c echo.Context) error {
	var forum entity.Forum
	c.Bind(&forum)

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	forum.ID = uuid

	data, err := fh.ForumU.Create(&forum)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(http.StatusBadRequest, "Failed create forums", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", data))
}

func (fh ForumHandler) Update(c echo.Context) error {
	forum := entity.Forum{}
	id := c.Param("id")
	c.Bind(&forum)
	data, err := fh.ForumU.Update(id, &forum)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(http.StatusBadRequest, "Failed update forums", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", data))
}

func (fh ForumHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := fh.ForumU.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(http.StatusBadRequest, "Failed get all forums", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData(http.StatusOK, "success", id))
}
