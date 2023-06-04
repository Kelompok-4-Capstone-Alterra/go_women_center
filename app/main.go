package main

import (
	"net/http"
	"os"

	CounselorAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/handler"
	CounselorAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/repository"
	CounselorAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/app/config"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	TopicHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/topic/handler"
	UserAuthHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/handler"
	CounselorUserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/handler"
	CounselorUserRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/repository"
	CounselorUserUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/usecase"
	ReviewUserRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/review/repository"
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

	sslconf := config.SSLconf{
		SSL_CERT:        os.Getenv("SSL_CERT"),
		SSL_PRIVATE_KEY: os.Getenv("SSL_PRIVATE_KEY"),
	}

	googleOauthConfig := &oauth2.Config{
		RedirectURL:  "http://localhost:8080/google/callback",
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	db := dbconf.InitDB()
	sslconf.InitSSL()
	// helper
	// googleUUID := helper.NewGoogleUUID()
	image := helper.NewImage("women-center")
	// log.Print(db, googleUUID)
	e := echo.New()

	// middleware 
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})
	
	
	// topic
	{
		topicHandler := TopicHandler.NewTopicHandler()
		e.GET("/topics", topicHandler.GetAll)
	}
	
	// admin
	groupAdmins := e.Group("/admin")
	{
		counselorRepo := CounselorAdminRepository.NewMysqlCounselorRepository(db)
		counselorUsecase := CounselorAdminUsecase.NewCounselorUsecase(counselorRepo, image)
		counselorHandler := CounselorAdminHandler.NewCounselorHandler(counselorUsecase)
		{
			groupAdmins.POST("/counselors", counselorHandler.Create)
			groupAdmins.GET("/counselors", counselorHandler.GetAll)
			groupAdmins.GET("/counselors/:id", counselorHandler.GetById)
			groupAdmins.PUT("/counselors/:id", counselorHandler.Update)
			groupAdmins.DELETE("/counselors/:id", counselorHandler.Delete)
		}
	}

	// users
	groupUsers := e.Group("/users", func(next echo.HandlerFunc) echo.HandlerFunc {
		// dummy user
		return func(c echo.Context) error {
			user := &entity.UserDecodeJWT{
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
	
	{
		userAuthHandler := UserAuthHandler.NewUserHandler(googleOauthConfig)
		{
			groupUsers.GET("/google/login", userAuthHandler.LoginHandler)
			groupUsers.GET("/google/callback", userAuthHandler.LoginGoogleCallback)
		}
		counselorRepo := CounselorUserRepository.NewMysqlCounselorRepository(db)
		reviewRepo := ReviewUserRepository.NewMysqlReviewRepository(db)
		counselorUsecase := CounselorUserUsecase.NewCounselorUsecase(counselorRepo, reviewRepo)
		counselorHandler := CounselorUserHandler.NewCounselorHandler(counselorUsecase)
		{	
			groupUsers.GET("/counselors", counselorHandler.GetAll)
			groupUsers.GET("/counselors/:id", counselorHandler.GetById)
		}
		groupUsers.POST("/counselors/:id/reviews", counselorHandler.CreateReview)
		groupUsers.GET("/counselors/:id/reviews", counselorHandler.GetAllReview)
	}

	// ssl
	// e.Logger.Fatal(e.StartTLS(":8080", "./ssl/certificate.crt", "./ssl/private.key"))

	e.Logger.Fatal(e.Start(":8080"))
}