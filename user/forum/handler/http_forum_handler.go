package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum"
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
	var getAllRequest forum.GetAllRequest
	c.Bind(&getAllRequest)

	user, ok := c.Get("user").(*helper.JwtCustomUserClaims)
	if !ok || user == nil {
		getAllRequest.UserId = ""
	} else {
		getAllRequest.UserId = user.ID
	}

	getAllRequest.Page, getAllRequest.Offset, getAllRequest.Limit = helper.GetPaginateData(getAllRequest.Page, getAllRequest.Limit)

	if err := isRequestValid(getAllRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	forums, totalPages, err := fh.ForumU.GetAll(getAllRequest)
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

func (fh ForumHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	var user_id string
	user, ok := c.Get("user").(*helper.JwtCustomUserClaims)
	if !ok || user == nil {
		user_id = ""
	} else {
		user_id = user.ID
	}

	forum, err := fh.ForumU.GetById(id, user_id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("success to get forum data details", http.StatusOK, forum))
}

func (fh ForumHandler) Create(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	var createRequest forum.CreateRequest
	c.Bind(&createRequest)

	uuidWithHyphen := uuid.New()
	createRequest.ID = uuidWithHyphen.String()
	createRequest.UserId = user.ID

	helper.RemoveWhiteSpace(&createRequest)

	if err := isRequestValid(createRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	err := fh.ForumU.Create(&createRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("successfully created forum data", http.StatusOK, nil))
}

func (fh ForumHandler) Update(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	id := c.Param("id")
	var updateRequest forum.UpdateRequest
	c.Bind(&updateRequest)

	helper.RemoveWhiteSpace(&updateRequest)

	if err := isRequestValid(updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	err := fh.ForumU.Update(id, user.ID, &updateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("successfully changed forum data", http.StatusOK, nil))
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
