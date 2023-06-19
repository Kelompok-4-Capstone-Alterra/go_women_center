package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	userForum "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserForumHandlerInterface interface {
	Create(c echo.Context) error
}

type UserForumHandler struct {
	UserForumU usecase.UserForumUsecaseInterface
}

func NewUserForumHandler(UserForumU usecase.UserForumUsecaseInterface) UserForumHandlerInterface {
	return &UserForumHandler{
		UserForumU: UserForumU,
	}
}

func (ufh UserForumHandler) Create(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	var createUserForum userForum.CreateRequest
	c.Bind(&createUserForum)
	uuidWithHyphen := uuid.New()
	createUserForum.ID = uuidWithHyphen.String()
	createUserForum.UserId = user.ID

	if err := isRequestValid(createUserForum); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}
	err := ufh.UserForumU.Create(&createUserForum)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(userForum.ErrFailedCreateReadingList.Error(), http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("Successfully joined the forum", http.StatusOK, nil))
}
