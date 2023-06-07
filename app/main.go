package main

import (
	"log"
	"net/http"
	"os"

	AdminAuthHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/handler"
	adminAuthMidd "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/handler/middleware"
	AdminAuthRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/repository"
	AdminAuthUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/usecase"
	CounselorAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/handler"
	CounselorAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/repository"
	CounselorAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/app/config"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	TopicHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/topic/handler"
	UserAuthHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/handler"
	userAuthMidd "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/handler/middleware"
	UserAuthRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/repository"
	UserAuthUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/usecase"
	CounselorUserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/handler"
	CounselorUserRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/repository"
	CounselorUserUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/usecase"
	ForumAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/handler"
	ForumAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/repository"
	ForumAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/usecase"
	ReviewUserRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/review/repository"
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

	// sslconf := config.SSLconf{
	// 	SSL_CERT:        os.Getenv("SSL_CERT"),
	// 	SSL_PRIVATE_KEY: os.Getenv("SSL_PRIVATE_KEY"),
	// }

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
	// sslconf.InitSSL()

	// helper
	jwtConf := helper.NewAuthJWT(os.Getenv("JWT_SECRET_USER"), os.Getenv("JWT_SECRET_ADMIN"))
	encryptor := helper.NewEncryptor()
	otpGenerator := helper.NewOtpGenerator()
	image := helper.NewImage("women-center")
	googleUUID := helper.NewGoogleUUID()
	log.Print(db, googleUUID)

	// handler
	userAuthRepo := UserAuthRepo.NewUserRepository(db)
	otpRepo := UserAuthRepo.NewLocalCache(config.CLEANUP_INTERVAL)
	userAuthUsecase := UserAuthUsecase.NewUserUsecase(userAuthRepo, googleUUID, &mailConf, otpRepo, otpGenerator, encryptor)
	userAuthHandler := UserAuthHandler.NewUserHandler(userAuthUsecase, googleOauthConfig, jwtConf)

	userCounselorRepo := CounselorUserRepository.NewMysqlCounselorRepository(db)
	userReviewRepo := ReviewUserRepository.NewMysqlReviewRepository(db)
	userCounselorUsecase := CounselorUserUsecase.NewCounselorUsecase(userCounselorRepo, userReviewRepo, userAuthRepo)
	userCounselorHandler := CounselorUserHandler.NewCounselorHandler(userCounselorUsecase)

	adminAuthRepo := AdminAuthRepo.NewAdminRepo(db)
	adminAuthUsecase := AdminAuthUsecase.NewAuthUsecase(adminAuthRepo, encryptor)
	adminAuthHandler := AdminAuthHandler.NewAuthHandler(adminAuthUsecase, jwtConf)

	adminCounselorRepo := CounselorAdminRepository.NewMysqlCounselorRepository(db)
	adminCounselorUsecase := CounselorAdminUsecase.NewCounselorUsecase(adminCounselorRepo, image)
	adminCounselorHandler := CounselorAdminHandler.NewCounselorHandler(adminCounselorUsecase)

	forumR := ForumAdminRepository.NewMysqlForumRepository(db)
	forumU := ForumAdminUsecase.NewForumUsecase(forumR)
	forumH := ForumAdminHandler.NewForumHandler(forumU)

	userForumR := UserForumAdminRepository.NewMysqlUserForumRepository(db)
	userForumU := UserForumAdminUsecase.NewUserForumUsecase(userForumR)
	userForumH := UserForumAdminHandler.NewUserForumHandler(userForumU)

	topicHandler := TopicHandler.NewTopicHandler()

	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})

	e.GET("/topics", topicHandler.GetAll)
	e.POST("/verify", userAuthHandler.VerifyEmailHandler)
	e.POST("/register", userAuthHandler.RegisterHandler)
	e.POST("/login", userAuthHandler.LoginHandler)
	e.GET("/google/login", userAuthHandler.LoginGoogleHandler)
	e.GET("/google/callback", userAuthHandler.LoginGoogleCallback)
	e.POST("/admin/login", adminAuthHandler.LoginHandler)

	users := e.Group("/users")
	{
		users.GET("/counselors", userCounselorHandler.GetAll)
	}

	// e.GET("/user/forums", forumH.GetAll)
	// e.GET("/user/forums/categories/:id", forumH.GetByCategory)
	// e.GET("/user/forums/my", forumH.GetByMyForum)
	// e.GET("/user/forums/:id", forumH.GetById)
	// e.POST("/user/forums", forumH.Create)
	// e.PUT("/user/forums/:id", forumH.Update)
	// e.DELETE("/user/forums/:id", forumH.Delete)
	// e.POST("/user/forums/joins", userForumH.Create)

	restrictUsers := e.Group("/users", userAuthMidd.JWTUser())
	{
		restrictUsers.GET("/profile", func(c echo.Context) error {
			user := c.Get("user").(*helper.JwtCustomUserClaims)
			return c.JSON(http.StatusOK, user)
		})

		restrictUsers.GET("/counselors/:id", userCounselorHandler.GetById)
		restrictUsers.POST("/counselors/:id/reviews", userCounselorHandler.CreateReview)
		restrictUsers.GET("/counselors/:id/reviews", userCounselorHandler.GetAllReview)

		restrictUsers.GET("/forums", forumH.GetAll)
		restrictUsers.GET("/forums/:id", forumH.GetById)
		restrictUsers.POST("/forums", forumH.Create)
		restrictUsers.PUT("/forums/:id", forumH.Update)
		restrictUsers.DELETE("/forums/:id", forumH.Delete)
		restrictUsers.POST("/forums/joins", userForumH.Create)
	}

	restrictAdmin := e.Group("/admin", adminAuthMidd.JWTAdmin())

	{
		restrictAdmin.GET("/profile", func(c echo.Context) error {
			admin := c.Get("admin").(*helper.JwtCustomAdminClaims)
			return c.JSON(http.StatusOK, admin)
		})

		restrictAdmin.GET("/counselors", adminCounselorHandler.GetAll)
		restrictAdmin.POST("/counselors", adminCounselorHandler.Create)
		restrictAdmin.GET("/counselors/:id", adminCounselorHandler.GetById)
		restrictAdmin.PUT("/counselors/:id", adminCounselorHandler.Update)
		restrictAdmin.DELETE("/counselors/:id", adminCounselorHandler.Delete)

	}
	e.Logger.Fatal(e.Start(":8080"))
	// ssl
	// e.Logger.Fatal(e.StartTLS(":8080", "./ssl/certificate.crt", "./ssl/private.key"))

	// e.Logger.Fatal(e.Start(":8080"))
}
