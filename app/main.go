package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/app/config"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	UserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/handler"
	UserRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/repository"
	UserUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/usecase"
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

	mailConf := helper.NewEmailSender(
		587,
		"smtp.gmail.com",
		os.Getenv("CONFIG_AUTH_EMAIL"),
		os.Getenv("CONFIG_AUTH_PASSWORD"),
		"Women Center <ivanhilmideran@gmail.com>", //TODO: set email to the proper one
	)

	db := dbconf.InitDB()
	googleUUID := helper.NewGoogleUUID()
	log.Print(db, googleUUID)

	userRepo := UserRepo.NewUserRepo(db)
	otpRepo := UserRepo.NewLocalCache(config.CLEANUP_INTERVAL)
	userUsecase := UserUsecase.NewUserUsecase(userRepo, googleUUID, &mailConf, otpRepo)
	userHandler := UserHandler.NewUserHandler(userUsecase, googleOauthConfig)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})

	e.POST("/verify", userHandler.VerifyEmailHandler)
	e.POST("/register", userHandler.RegisterHandler)
	e.POST("/login", userHandler.LoginHandler)
	e.GET("/google/login", userHandler.LoginGoogleHandler)
	e.GET("/google/callback", userHandler.LoginGoogleCallback)

	e.Logger.Fatal(e.Start(":8080"))
}
