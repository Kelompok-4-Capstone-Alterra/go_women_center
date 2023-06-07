package handler

import (
	"net/http"

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
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	getCreated := c.QueryParam("created")
	getTopic := c.QueryParam("topic")
	getPopular := c.QueryParam("popular")
	getCategories := c.QueryParam("categories")
	getMyForum := c.QueryParam("myforum")

	forums, err := fh.ForumU.GetAll(user.ID, getTopic, getPopular, getCreated, getCategories, getMyForum)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("Success to get all forums data", http.StatusOK, forums))
}

func (fh ForumHandler) GetById(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	id := c.Param("id")
	forum, err := fh.ForumU.GetById(id, user.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("Success to get detail forum data", http.StatusOK, forum))
}

func (fh ForumHandler) Create(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	var forum entity.Forum
	c.Bind(&forum)

	uuidWithHyphen := uuid.New()
	forum.ID = uuidWithHyphen.String()
	forum.UserId = user.ID

	err := fh.ForumU.Create(&forum)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("Successfully created forum data", http.StatusOK, nil))
}

func (fh ForumHandler) Update(c echo.Context) error {
	forum := entity.Forum{}
	id := c.Param("id")
	c.Bind(&forum)
	err := fh.ForumU.Update(id, &forum)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData("Failed update forums", http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("Successfully updated forum data", http.StatusOK, nil))
}

func (fh ForumHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := fh.ForumU.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData("Failed get all forums", http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("Successfully deleted forum data", http.StatusOK, nil))
}