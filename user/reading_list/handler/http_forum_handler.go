package handler

import (
	"net/http"
	"strconv"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ReadingListHandlerInterface interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type ReadingListHandler struct {
	ReadingListU usecase.ReadingListUsecaseInterface
}

func NewReadingListHandler(ReadingListU usecase.ReadingListUsecaseInterface) ReadingListHandlerInterface {
	return &ReadingListHandler{
		ReadingListU: ReadingListU,
	}
}

func (rlh ReadingListHandler) GetAll(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	getName := c.QueryParam("name")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	page, offset, limit := helper.GetPaginateData(page, limit)
	reading_list, totalPages, err := rlh.ReadingListU.GetAll(user.ID, getName, offset, limit)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("success to get all reading list data", http.StatusOK, echo.Map{
		"reading_list":  reading_list,
		"current_pages": page,
		"total_pages":   totalPages,
	}))
}

func (rlh ReadingListHandler) GetById(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	id := c.Param("id")
	reading_list, err := rlh.ReadingListU.GetById(id, user.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("success to get reading list data details", http.StatusOK, reading_list))
}

func (rlh ReadingListHandler) Create(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	var reading_list entity.ReadingList
	c.Bind(&reading_list)

	uuidWithHyphen := uuid.New()
	reading_list.ID = uuidWithHyphen.String()
	reading_list.UserId = user.ID

	err := rlh.ReadingListU.Create(&reading_list)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("successfully created reading list data", http.StatusOK, nil))
}

func (rlh ReadingListHandler) Update(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	reading_list := entity.ReadingList{}
	id := c.Param("id")
	c.Bind(&reading_list)
	err := rlh.ReadingListU.Update(id, user.ID, &reading_list)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("successfully changed reading list data", http.StatusOK, nil))
}

func (rlh ReadingListHandler) Delete(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	id := c.Param("id")
	err := rlh.ReadingListU.Delete(id, user.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("successfully deleted reading list data", http.StatusOK, nil))
}
