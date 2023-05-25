package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/app/config"
	CounselorHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/counselor/handler"
	CounselorRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/counselor/repository"
	CounselorUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/counselor/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	UserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/handler"
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

	googleOauthConfig := &oauth2.Config{
		RedirectURL:  "http://localhost:8080/google/callback",
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	db := dbconf.InitDB()
	googleUUID := helper.NewGoogleUUID()
	log.Print(db, googleUUID)

	userHandler := UserHandler.NewUserHandler(googleOauthConfig)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	
	
	// e.Static("/images", "images/")

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})

	// user
	e.GET("/google/login", userHandler.LoginHandler)
	e.GET("/google/callback", userHandler.LoginGoogleCallback)


	// counselor

	counselorRepo := CounselorRepository.NewMysqlCounselorRepository(db)
	counselorUsecase := CounselorUsecase.NewCounselorUsecase(counselorRepo)
	counselorHandler := CounselorHandler.NewCounselorHandler(counselorUsecase)

	{
		e.POST("/admins/counselors", counselorHandler.Create)
	}
	
	e.GET("/testimage", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<img src='https://bucket-test-556.s3.ap-southeast-1.amazonaws.com/644b12a8-fafd-11ed-95df-5efc22537c1d_status_pending.png'/>")
	})

	e.Logger.Fatal(e.Start(":8080"))
}