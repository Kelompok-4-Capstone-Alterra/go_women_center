package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	request "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
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
	var getAllParam request.QueryParamRequest
	c.Bind(&getAllParam)
	getAllParam.IdUser = user.ID
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
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	id := c.Param("id")
	forum, err := fh.ForumU.GetById(id, user.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("success to get forum data details", http.StatusOK, forum))
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
	return c.JSON(http.StatusOK, helper.ResponseData("successfully created forum data", http.StatusOK, nil))
}

func (fh ForumHandler) Update(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	forum := entity.Forum{}
	id := c.Param("id")
	c.Bind(&forum)
	err := fh.ForumU.Update(id, user.ID, &forum)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("successfully updated forum data", http.StatusOK, nil))
}

func (fh ForumHandler) Delete(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	id := c.Param("id")
	err := fh.ForumU.Delete(id, user.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("successfully deleted forum data", http.StatusOK, nil))
}
