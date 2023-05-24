package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type userHandler struct {
	OauthConf *oauth2.Config
}

func NewUserHandler(oauthConf *oauth2.Config) *userHandler {
	return &userHandler{
		OauthConf: oauthConf,
	}
}

var oauthstatemap = map[string]bool{}

type userInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

func (u *userHandler) LoginHandler(e echo.Context) error {
	g := helper.NewGoogleUUID()
	uuid, _ := g.GenerateUUID()
	// oauthStateString = uuid
	oauthstatemap[uuid] = true

	url := u.OauthConf.AuthCodeURL(uuid)

	return e.Redirect(http.StatusTemporaryRedirect, url)
}

func (u *userHandler) LoginGoogleCallback(e echo.Context) error {
	content, err := u.getUserInfo(e.FormValue("state"), e.FormValue("code"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	//TODO: send token to response
	return e.JSON(http.StatusOK, content)
}

func (u *userHandler) getUserInfo(state, code string) (userInfo, error) {

	UserInfo := userInfo{}
	if !oauthstatemap[state] {
		return UserInfo, fmt.Errorf("invalid oauth state")
	}
	defer delete(oauthstatemap, state)

	token, err := u.OauthConf.Exchange(context.Background(), code)
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
