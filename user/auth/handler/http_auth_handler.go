package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	user "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/usecase"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type userHandler struct {
	Usecase   usecase.UserUsecase
	OauthConf *oauth2.Config
	JwtConf   helper.AuthJWT
	Validator helper.Validator
}

func NewUserHandler(u usecase.UserUsecase, oauthConf *oauth2.Config, jwtConf helper.AuthJWT, vld helper.Validator) *userHandler {
	return &userHandler{
		Usecase:   u,
		OauthConf: oauthConf,
		JwtConf:   jwtConf,
		Validator: vld,
	}
}

var oauthstatemap = map[string]bool{}

func (h *userHandler) LoginGoogleHandler(c echo.Context) error {
	g := helper.NewGoogleUUID()
	uuid, _ := g.GenerateUUID()
	// oauthStateString = uuid
	oauthstatemap[uuid] = true

	url := h.OauthConf.AuthCodeURL(uuid)

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *userHandler) LoginGoogleCallback(c echo.Context) error {
	content, err := h.getUserInfo(c.FormValue("state"), c.FormValue("code"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			http.StatusInternalServerError,
			err.Error(),
			nil,
		))
	}

	//TODO: send token to response

	return c.JSON(http.StatusOK, content)
}

func (h *userHandler) getUserInfo(state, code string) (user.UserOauthInfo, error) {

	UserInfo := user.UserOauthInfo{}
	if !oauthstatemap[state] {
		return UserInfo, fmt.Errorf("invalid oauth state")
	}
	defer delete(oauthstatemap, state)

	token, err := h.OauthConf.Exchange(context.Background(), code)
	if err != nil {
		return UserInfo, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return UserInfo, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&UserInfo)
	if err != nil {
		return UserInfo, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return UserInfo, nil
}

func (h *userHandler) VerifyEmailHandler(c echo.Context) error { // TODO: rename with suffix handler
	emailDTO := user.VerifyEmailRequest{}
	err := c.Bind(&emailDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			http.StatusInternalServerError,
			err.Error(),
			nil,
		))
	}

	err = h.Validator.ValidateStruct(emailDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			http.StatusBadRequest,
			err.Error(),
			nil,
		))
	}

	err = h.Usecase.VerifyEmail(emailDTO.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			http.StatusInternalServerError,
			err.Error(), //TODO: write better error message
			nil,
		))
	}

	return c.JSON(http.StatusOK, helper.ResponseData(
		http.StatusOK,
		"success sending otp, valid for 1 minute",
		nil,
	))
}

func (h *userHandler) RegisterHandler(c echo.Context) error {
	reqDTO := user.RegisterUserRequest{}
	err := c.Bind(&reqDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			http.StatusInternalServerError,
			err.Error(),
			nil,
		))
	}

	err = h.Validator.ValidateStruct(reqDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			http.StatusBadRequest,
			err.Error(),
			nil,
		))
	}

	err = h.Usecase.Register(reqDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			http.StatusInternalServerError,
			err.Error(),
			nil,
		))
	}

	//TODO: send token to response
	return c.JSON(http.StatusOK, helper.ResponseData(
		http.StatusOK,
		"register success",
		nil,
	))
}

func (h *userHandler) LoginHandler(c echo.Context) error {
	reqDTO := user.LoginUserRequest{}
	err := c.Bind(&reqDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			http.StatusInternalServerError,
			err.Error(),
			nil,
		))
	}

	err = h.Validator.ValidateStruct(reqDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			http.StatusBadRequest,
			err.Error(),
			nil,
		))
	}

	data, err := h.Usecase.Login(reqDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			http.StatusInternalServerError,
			err.Error(),
			nil,
		))
	}

	token, err := h.JwtConf.GenerateUserToken(data.ID, data.Email, constant.Auth)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			http.StatusInternalServerError,
			"fail to generate user token",
			nil,
		))
	}

	return c.JSON(http.StatusOK, helper.ResponseData(
		http.StatusOK,
		"login success",
		map[string]interface{}{
			"token": token,
		},
	))
}
