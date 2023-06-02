package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/app/config"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	UserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/handler"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	CareerAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/career/handler"
	CareerAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/career/repository"
	CareerAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/career/usecase"
	
	CareerUserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career/handler"
	CareerUserRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career/repository"
	CareerUserUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career/usecase"
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

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})

	e.GET("/google/login", userHandler.LoginHandler)
	e.GET("/google/callback", userHandler.LoginGoogleCallback)

	groupAdmins := e.Group("/admin")
	{
		careerRepo := CareerAdminRepository.NewMysqlCareerRepository(db)
		careerUsecase := CareerAdminUsecase.NewCareerUsecase(careerRepo)
		careerHandler := CareerAdminHandler.NewCareerHandler(careerUsecase)
		{
			groupAdmins.POST("/careers", careerHandler.Create)
			groupAdmins.GET("/careers", careerHandler.GetAll)
			groupAdmins.GET("/careers/:id", careerHandler.GetById)
			groupAdmins.PUT("/careers/:id", careerHandler.Update)
			groupAdmins.DELETE("/careers/:id", careerHandler.Delete)
		}
	}

	groupUsers := e.Group("/users")
	{
		careerUserRepo := CareerUserRepository.NewMysqlCareerRepository(db)
		careerUsecase := CareerUserUsecase.NewCareerUsecase(careerUserRepo)
		careerHandler := CareerUserHandler.NewCareerHandler(careerUsecase)
		{
			groupUsers.GET("/careers", careerHandler.GetAll)
		}
	}
	

	e.Logger.Fatal(e.Start(":8080"))
}
