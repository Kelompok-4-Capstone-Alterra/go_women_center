package handler

import (
	"net/http"
	"strconv"

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
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	page, offset, limit := helper.GetPaginateData(page, limit)

	forums, totalPages, err := fh.ForumU.GetAll(user.ID, getTopic, getPopular, getCreated, getCategories, getMyForum, offset, limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData("Failed to get all forum data", http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("Success to get all forum data", http.StatusOK, echo.Map{
		"forums":        forums,
		"current_pages": page,
		"total_pages":   totalPages,
	}))
}

func (fh ForumHandler) GetById(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	id := c.Param("id")
	forum, err := fh.ForumU.GetById(id, user.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData("Failed to get forum data details", http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("Success to get forum data details", http.StatusOK, forum))
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
		return c.JSON(http.StatusBadRequest, helper.ResponseData("Failed created forum data", http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("Successfully created forum data", http.StatusOK, nil))
}

func (fh ForumHandler) Update(c echo.Context) error {
	forum := entity.Forum{}
	id := c.Param("id")
	c.Bind(&forum)
	err := fh.ForumU.Update(id, &forum)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData("Failed to change forum data", http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("Successfully changed forum data", http.StatusOK, nil))
}

func (fh ForumHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := fh.ForumU.Delete(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData("Failed to delete forum data", http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("Successfully deleted forum data", http.StatusOK, nil))
}
