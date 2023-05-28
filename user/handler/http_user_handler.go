package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/usecase"
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
		JwtConf: jwtConf,
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
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
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
	emailDTO := user.VerifyEmailDTO{}
	err := c.Bind(&emailDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	err = h.Usecase.VerifyEmail(emailDTO.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(), //TODO: write better error message
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success sending otp",
	})
}

func (h *userHandler) RegisterHandler(c echo.Context) error {
	reqDTO := user.RegisterUserDTO{}
	err := c.Bind(&reqDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	// TODO: validate req

	err = h.Usecase.Register(reqDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	//TODO: send token to response
	return c.JSON(http.StatusOK, echo.Map{
		"message": "register success",
	})
}

func (h *userHandler) LoginHandler(c echo.Context) error {
	reqDTO := user.LoginUserDTO{}
	err := c.Bind(&reqDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	// TODO: validate req

	data, err := h.Usecase.Login(reqDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	token, err := h.JwtConf.GenerateToken(data.ID, data.Email, constant.Auth)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "login success",
		"token": token,
	})
}
