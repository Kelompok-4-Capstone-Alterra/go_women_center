package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	readingList "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list"
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
	var getAllParams readingList.GetAllRequest

	c.Bind(&getAllParams)

	if err := isRequestValid(getAllParams); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	getAllParams.UserId = user.ID

	getAllParams.Page, getAllParams.Offset, getAllParams.Limit = helper.GetPaginateData(getAllParams.Page, getAllParams.Limit)

	reading_list, totalPages, err := rlh.ReadingListU.GetAll(getAllParams)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	if getAllParams.Page > totalPages {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(readingList.ErrPageNotFound.Error(), http.StatusBadRequest, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success to get all reading list data", http.StatusOK, echo.Map{
		"reading_list":  reading_list,
		"current_pages": getAllParams.Page,
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
	var createForum readingList.CreateRequest
	c.Bind(&createForum)

	uuidWithHyphen := uuid.New()
	createForum.ID = uuidWithHyphen.String()
	createForum.UserId = user.ID

	helper.RemoveWhiteSpace(&createForum)

	if err := isRequestValid(createForum); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	err := rlh.ReadingListU.Create(&createForum)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("successfully created reading list data", http.StatusOK, nil))
}

func (rlh ReadingListHandler) Update(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	var updateRequest readingList.UpdateRequest
	id := c.Param("id")
	c.Bind(&updateRequest)

	helper.RemoveWhiteSpace(&updateRequest)

	if err := isRequestValid(updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	err := rlh.ReadingListU.Update(id, user.ID, &updateRequest)

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
