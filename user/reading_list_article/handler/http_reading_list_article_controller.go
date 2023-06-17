package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	readingListArticle "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list_article"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list_article/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ReadingListArticleHandlerInterface interface {
	Create(c echo.Context) error
	Delete(c echo.Context) error
}

type ReadingListArticleHandler struct {
	ReadingListArticleU usecase.ReadingListArticleUsecaseInterface
}

func NewReadingListArticleHandler(ReadingListArticleU usecase.ReadingListArticleUsecaseInterface) ReadingListArticleHandlerInterface {
	return &ReadingListArticleHandler{
		ReadingListArticleU: ReadingListArticleU,
	}
}

func (rlah ReadingListArticleHandler) Create(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	var createRequest readingListArticle.CreateRequest
	c.Bind(&createRequest)

	uuidWithHyphen := uuid.New()
	createRequest.ID = uuidWithHyphen.String()
	createRequest.UserId = user.ID
	err := rlah.ReadingListArticleU.Create(&createRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("Successfully joined the forum", http.StatusOK, nil))
}

func (rlah ReadingListArticleHandler) Delete(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	id := c.Param("id")
	err := rlah.ReadingListArticleU.Delete(id, user.ID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("successfully deleted reading list data", http.StatusOK, nil))
}
