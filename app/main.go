package main

import (
	"net/http"
	"os"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/app/config"
	CounselorHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/counselor/handler"
	CounselorRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/counselor/repository"
	CounselorUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/counselor/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
	ReviewRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/review/repository"
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
	// googleUUID := helper.NewGoogleUUID()
	// log.Print(db, googleUUID)

	userHandler := UserHandler.NewUserHandler(googleOauthConfig)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})

	// user
	e.GET("/google/login", userHandler.LoginHandler)
	e.GET("/google/callback", userHandler.LoginGoogleCallback)


	groupAdmins := e.Group("/admins")
	groupUsers := e.Group("/users", func(next echo.HandlerFunc) echo.HandlerFunc {
		// dummy user
		return func(c echo.Context) error {
			user := &domain.UserDecodeJWT{
				ID: "05b9B469-fc5d-21ed-ad1c-5efc22537c1d",
				Name: "dummy",
				Email: "dummy@gmail.com",
				Method: "basic",
				Role: "user",
			}
			c.Set("user", user)

			return next(c)
		}
	})

	// counselor

	counselorRepo := CounselorRepository.NewMysqlCounselorRepository(db)
	reviewRepo := ReviewRepository.NewMysqlReviewRepository(db)
	counselorUsecase := CounselorUsecase.NewCounselorUsecase(counselorRepo, reviewRepo)
	counselorHandler := CounselorHandler.NewCounselorHandler(counselorUsecase)

	// admins
	{
		groupAdmins.POST("/counselors", counselorHandler.Create)
		groupAdmins.GET("/counselors", counselorHandler.GetAll)
		groupAdmins.GET("/counselors/:id", counselorHandler.GetById)
		groupAdmins.PUT("/counselors/:id", counselorHandler.Update)
		groupAdmins.DELETE("/counselors/:id", counselorHandler.Delete)
	}

	// users
	{
		groupUsers.GET("/counselors", counselorHandler.GetAll)
		groupUsers.GET("/counselors/:id", counselorHandler.GetById)
		groupUsers.POST("/counselors/:id/reviews", counselorHandler.CreateReview)
		// groupUsers.GET("/counselors/:id/reviews", counselorHandler.GetAllReview)
	}

	e.Logger.Fatal(e.Start(":8080"))
}