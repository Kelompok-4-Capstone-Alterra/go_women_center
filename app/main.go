package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/app/config"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	dbconf := config.DBconf{
		DB_Username: os.Getenv("DB_USERNAME"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Name:     os.Getenv("DB_NAME"),
	}
	db := dbconf.InitDB()
	googleUUID := helper.NewGoogleUUID()
	log.Print(db, googleUUID)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})

	e.GET("/google/login", LoginHandler)
	e.GET("/google/callback", LoginGoogleCallback)
	e.GET("/login", LoginViewHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/google/callback",
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func LoginViewHandler(e echo.Context) error {
	return e.HTML(http.StatusOK, "<a href='/google/login'>Sign in with Google</a>")
}

func LoginGoogleCallback(e echo.Context) error {
	content, err := getUserInfo(e.FormValue("state"), e.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		return e.Redirect(http.StatusTemporaryRedirect, "/")
	}
	
	return e.JSON(http.StatusOK, content)	
}

var oauthStateString string
type userInfo struct{
	ID string `json:"id"`
	Email string `json:"email"`
	VerifiedEmail bool `json:"verified_email"`
	Picture string `json:"picture"`
}

func getUserInfo(state, code string) (userInfo, error) {

	UserInfo := userInfo{}

	if state != oauthStateString {
		return UserInfo, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)

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

func LoginHandler(e echo.Context) error{

	g := helper.NewGoogleUUID()

	uuid, _ := g.GenerateUUID()

	oauthStateString = uuid

	url := googleOauthConfig.AuthCodeURL(uuid)

	return e.Redirect(http.StatusTemporaryRedirect, url)
}
