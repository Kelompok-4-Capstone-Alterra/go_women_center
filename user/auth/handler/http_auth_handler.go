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
}

func NewUserHandler(u usecase.UserUsecase, oauthConf *oauth2.Config, jwtConf helper.AuthJWT) *userHandler {
	return &userHandler{
		Usecase:   u,
		OauthConf: oauthConf,
		JwtConf:   jwtConf,
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
			err.Error(),
			http.StatusInternalServerError,
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

func (h *userHandler) VerifyUniqueCredential(c echo.Context) error {
	verifyRequest := user.VerifyUniqueCredentialRequest{}
	err := c.Bind(&verifyRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	if err := isRequestValid(verifyRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	err = h.Usecase.CheckUnique(verifyRequest.Email, verifyRequest.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	return c.JSON(http.StatusOK, helper.ResponseData(
		"credential is available",
		http.StatusOK,
		nil,
	))
}

func (h *userHandler) VerifyEmailHandler(c echo.Context) error { // TODO: rename with suffix handler
	emailRequest := user.VerifyEmailRequest{}
	err := c.Bind(&emailRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	if err := isRequestValid(emailRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	err = h.Usecase.VerifyEmail(emailRequest.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(), //TODO: write better error message
			http.StatusInternalServerError,
			nil,
		))
	}

	return c.JSON(http.StatusOK, helper.ResponseData(
		"success sending otp, valid for 1 minute",
		http.StatusOK,
		nil,
	))
}

func (h *userHandler) RegisterHandler(c echo.Context) error {
	registerRequest := user.RegisterUserRequest{}
	err := c.Bind(&registerRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	if err := isRequestValid(registerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	err = h.Usecase.Register(registerRequest)
	if err != nil {

		status := http.StatusBadRequest

		switch err {
		case user.ErrUserIsRegistered:
			status = http.StatusConflict
		case user.ErrInternalServerError:
			status = http.StatusInternalServerError
		}

		return c.JSON(status, helper.ResponseData(
			err.Error(),
			status,
			nil,
		))
	}

	//TODO: send token to response
	return c.JSON(http.StatusOK, helper.ResponseData(
		"register success",
		http.StatusOK,
		nil,
	))
}

func (h *userHandler) LoginHandler(c echo.Context) error {
	loginRequest := user.LoginUserRequest{}
	err := c.Bind(&loginRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	if err := isRequestValid(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	data, err := h.Usecase.Login(loginRequest)
	if err != nil {

		status := http.StatusBadRequest

		switch err {
		case user.ErrInternalServerError:
			status = http.StatusInternalServerError
		}

		return c.JSON(status, helper.ResponseData(
			err.Error(),
			status,
			nil,
		))
	}
	fmt.Println(data)
	token, _ := h.JwtConf.GenerateUserToken(data.ID, data.Email, data.Username, constant.Auth)

	return c.JSON(http.StatusOK, helper.ResponseData(
		"login success",
		http.StatusOK,
		echo.Map{
			"token": token,
		},
	))
}

// forget pass feature handler for email verif
func (h *userHandler) CheckIsRegistered(c echo.Context) error {
	registerReq := user.VerifyEmailRequest{}
	err := c.Bind(&registerReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	if err := isRequestValid(registerReq); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	err = h.Usecase.CheckIsRegistered(registerReq.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	err = h.Usecase.VerifyEmail(registerReq.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	return c.JSON(http.StatusOK, helper.ResponseData(
		"success sending otp, valid for 1 minute",
		http.StatusOK,
		nil,
	))
}

// check otp then send generated new pass to user email
func (h *userHandler) ForgetPassword(c echo.Context) error {
	forgetPassReq := user.ForgetPasswordRequest{}
	err := c.Bind(&forgetPassReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	if err := isRequestValid(forgetPassReq); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(
			err.Error(),
			http.StatusBadRequest,
			nil,
		))
	}

	err = h.Usecase.ForgetPassword(forgetPassReq.Email, forgetPassReq.OTP)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(
			err.Error(),
			http.StatusInternalServerError,
			nil,
		))
	}

	return c.JSON(http.StatusOK, helper.ResponseData(
		"success sending new password",
		http.StatusOK,
		nil,
	))
}
