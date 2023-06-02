package main

import (
	"log"
	"os"

	AdminAuthHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/handler"
	AdminAuthRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/repository"
	AdminAuthUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/app/config"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	UserAuthHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/handler"
	UserAuthRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/repository"
	UserAuthUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/usecase"
	_ "github.com/joho/godotenv/autoload"
	echojwt "github.com/labstack/echo-jwt/v4"
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

	jwtConf := helper.NewAuthJWT(os.Getenv("JWT_SECRET"))

	db := dbconf.InitDB()
	googleUUID := helper.NewGoogleUUID()
	log.Print(db, googleUUID)

	userAuthRepo := UserAuthRepo.NewUserRepo(db)
	otpRepo := UserAuthRepo.NewLocalCache(config.CLEANUP_INTERVAL)
	userAuthUsecase := UserAuthUsecase.NewUserUsecase(userAuthRepo, googleUUID, &mailConf, otpRepo)
	userAuthHandler := UserAuthHandler.NewUserHandler(userAuthUsecase, googleOauthConfig, jwtConf)

	adminAuthRepo := AdminAuthRepo.NewAdminRepo(db)
	adminAuthUsecase := AdminAuthUsecase.NewAuthUsecase(adminAuthRepo)
	adminAuthHandler := AdminAuthHandler.NewAuthHandler(adminAuthUsecase, jwtConf)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	restricted := e.Group("/user")
	restricted.Use(echojwt.JWT([]byte(jwtConf.GetSecret())))

	e.POST("/verify", userAuthHandler.VerifyEmailHandler)
	e.POST("/register", userAuthHandler.RegisterHandler)
	e.POST("/login", userAuthHandler.LoginHandler)
	e.GET("/google/login", userAuthHandler.LoginGoogleHandler)
	e.GET("/google/callback", userAuthHandler.LoginGoogleCallback)

	e.POST("/admin/login", adminAuthHandler.LoginHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
