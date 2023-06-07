package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/usecase"
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

func (fh UserForumHandler) Create(c echo.Context) error {
	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	var userForum entity.UserForum
	c.Bind(&userForum)
	userForum.UserId = user.ID

	err := fh.UserForumU.Create(&userForum)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData("Failed to join forum", http.StatusBadRequest, nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseData("Success to join forum", http.StatusOK, nil))
}
