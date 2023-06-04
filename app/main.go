package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/app/config"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	ForumAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/handler"
	ForumAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/repository"
	ForumAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/usecase"
	UserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/handler"
	UserForumAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/handler"
	UserForumAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/repository"
	UserForumAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/usecase"
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

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})

	e.GET("/google/login", userHandler.LoginHandler)
	e.GET("/google/callback", userHandler.LoginGoogleCallback)

	forumR := ForumAdminRepository.NewMysqlForumRepository(db)
	forumU := ForumAdminUsecase.NewForumUsecase(forumR)
	forumH := ForumAdminHandler.NewForumHandler(forumU)
	e.GET("/user/forums", forumH.GetAll)
	e.GET("/user/forums/sort", forumH.GetAllSortBy)
	e.GET("/user/forums/categories/:id", forumH.GetByCategory)
	e.GET("/user/forums/my", forumH.GetByMyForum)
	e.GET("/user/forums/:id", forumH.GetById)
	e.POST("/user/forums", forumH.Create)
	e.PUT("/user/forums/:id", forumH.Update)
	e.DELETE("/user/forums/:id", forumH.Delete)

	// create user forum
	userForumR := UserForumAdminRepository.NewMysqlUserForumRepository(db)
	userForumU := UserForumAdminUsecase.NewUserForumUsecase(userForumR)
	userForumH := UserForumAdminHandler.NewUserForumHandler(userForumU)
	e.POST("/user/forums/joins", userForumH.Create)

	e.Logger.Fatal(e.Start(":8080"))
}
