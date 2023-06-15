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
	CounselorAdminRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/repository"
	CounselorAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/usecase"
	ForumAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum/handler"
	ForumAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum/repository"
	ForumAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum/usecase"
	AdminScheduleHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/schedule/handler"
	AdminScheduleRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/schedule/repository"
	AdminScheduleUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/schedule/usecase"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/app/config"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	TopicHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/topic/handler"
	UserAuthHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/handler"
	userAuthMidd "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/handler/middleware"
	UserAuthRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/repository"
	UserAuthUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/usecase"
	CounselorUserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/handler"
	CounselorUserRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/repository"
	CounselorUserUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/usecase"

	UserProfileHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/profile/handler"
	UserProfileRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/profile/repository"
	UserProfileUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/profile/usecase"

	ForumUserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/handler"
	ForumUserRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/repository"
	ForumUserUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/usecase"
	UserForumAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/handler"
	UserForumAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/repository"
	UserForumAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/usecase"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	UsersAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/users/handler"
	UsersAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/users/repository"
	UsersAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/users/usecase"

	CareerAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/career/handler"
	CareerAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/career/repository"
	CareerAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/career/usecase"

	CareerUserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career/handler"
	CareerUserRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career/repository"
	CareerUserUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/career/usecase"

	ReadingListHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list/handler"
	ReadingListRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list/repository"
	ReadingListUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list/usecase"

	ReadingListArticleHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list_article/handler"
	ReadingListArticleRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list_article/repository"
	ReadingListArticleUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/reading_list_article/usecase"
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

	userCounselorRepo := CounselorUserRepo.NewMysqlCounselorRepository(db)
	userReviewRepo := CounselorUserRepo.NewMysqlReviewRepository(db)

	userCounselorUsecase := CounselorUserUsecase.NewCounselorUsecase(userCounselorRepo, userReviewRepo, userAuthRepo)
	userCounselorHandler := CounselorUserHandler.NewCounselorHandler(userCounselorUsecase)

	userRepo := UserProfileRepo.NewMysqlUserRepository(db)
	userUsecase := UserProfileUsecase.NewProfileUsecase(userRepo, image, encryptor)
	userHandler := UserProfileHandler.NewProfileHandler(userUsecase)

	userCareerRepo := CareerUserRepository.NewMysqlCareerRepository(db)
	userCareerUsecase := CareerUserUsecase.NewCareerUsecase(userCareerRepo)
	userCareerHandler := CareerUserHandler.NewCareerHandler(userCareerUsecase)

	adminAuthRepo := AdminAuthRepo.NewAdminRepo(db)
	adminAuthUsecase := AdminAuthUsecase.NewAuthUsecase(adminAuthRepo, encryptor)
	adminAuthHandler := AdminAuthHandler.NewAuthHandler(adminAuthUsecase, jwtConf)

	adminCounselorRepo := CounselorAdminRepo.NewMysqlCounselorRepository(db)
	adminCounselorUsecase := CounselorAdminUsecase.NewCounselorUsecase(adminCounselorRepo, image)
	adminCounselorHandler := CounselorAdminHandler.NewCounselorHandler(adminCounselorUsecase)

	adminCareerRepo := CareerAdminRepository.NewMysqlCareerRepository(db)
	adminCareerUsecase := CareerAdminUsecase.NewCareerUsecase(adminCareerRepo, image)
	adminCareerHandler := CareerAdminHandler.NewCareerHandler(adminCareerUsecase)

	adminUsersRepo := UsersAdminRepository.NewMysqlUserRepository(db)
	adminUsersUsecase := UsersAdminUsecase.NewUserUsecase(adminUsersRepo)
	adminUsersHandler := UsersAdminHandler.NewUserHandler(adminUsersUsecase)

	adminScheduleRepo := AdminScheduleRepo.NewMysqlScheduleRepository(db)
	adminScheduleUsecase := AdminScheduleUsecase.NewScheduleUsecase(adminCounselorRepo, adminScheduleRepo, googleUUID)
	adminScheduleHandler := AdminScheduleHandler.NewScheduleHandler(adminScheduleUsecase)

	forumR := ForumUserRepository.NewMysqlForumRepository(db)
	forumU := ForumUserUsecase.NewForumUsecase(forumR)
	forumH := ForumUserHandler.NewForumHandler(forumU)

	forumAdminR := ForumAdminRepository.NewMysqlForumRepository(db)
	forumAdminU := ForumAdminUsecase.NewForumUsecase(forumAdminR)
	forumAdminH := ForumAdminHandler.NewForumHandler(forumAdminU)

	userForumR := UserForumAdminRepository.NewMysqlUserForumRepository(db)
	userForumU := UserForumAdminUsecase.NewUserForumUsecase(userForumR)
	userForumH := UserForumAdminHandler.NewUserForumHandler(userForumU)

	ReadingListR := ReadingListRepository.NewMysqlReadingListRepository(db)
	ReadingListU := ReadingListUsecase.NewReadingListUsecase(ReadingListR)
	ReadingListH := ReadingListHandler.NewReadingListHandler(ReadingListU)

	ReadingListArticleR := ReadingListArticleRepository.NewMysqlReadingListArticleRepository(db)
	ReadingListArticleU := ReadingListArticleUsecase.NewReadingListArticleUsecase(ReadingListArticleR)
	ReadingListArticleH := ReadingListArticleHandler.NewReadingListArticleHandler(ReadingListArticleU)

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
		users.GET("/careers", userCareerHandler.GetAll)
	}

	restrictUsers := e.Group("/users", userAuthMidd.JWTUser(), userAuthMidd.CheckUser(userAuthUsecase))

	{
		restrictUsers.GET("/profile", userHandler.GetById)
		restrictUsers.PUT("/profile", userHandler.Update)
		restrictUsers.PUT("/profile/password", userHandler.UpdatePassword)
		restrictUsers.GET("/counselors/:id", userCounselorHandler.GetById)
		restrictUsers.POST("/counselors/:id/reviews", userCounselorHandler.CreateReview)
		restrictUsers.GET("/counselors/:id/reviews", userCounselorHandler.GetAllReview)

		restrictUsers.GET("/forums", forumH.GetAll)
		restrictUsers.GET("/forums/:id", forumH.GetById)
		restrictUsers.POST("/forums", forumH.Create)
		restrictUsers.PUT("/forums/:id", forumH.Update)
		restrictUsers.DELETE("/forums/:id", forumH.Delete)
		restrictUsers.POST("/forums/joins", userForumH.Create)
		restrictUsers.GET("/careers/:id", userCareerHandler.GetById)

		restrictUsers.GET("/reading-lists", ReadingListH.GetAll)
		restrictUsers.GET("/reading-lists/:id", ReadingListH.GetById)
		restrictUsers.POST("/reading-lists", ReadingListH.Create)
		restrictUsers.PUT("/reading-lists/:id", ReadingListH.Update)
		restrictUsers.DELETE("/reading-lists/:id", ReadingListH.Delete)

		restrictUsers.POST("/reading-lists/save", ReadingListArticleH.Create)
		restrictUsers.DELETE("/reading-lists/save/:id", ReadingListArticleH.Delete)

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

		restrictAdmin.POST("/counselors/:id/schedules", adminScheduleHandler.Create)
		restrictAdmin.GET("/counselors/:id/schedules", adminScheduleHandler.GetByCounselorId)
		restrictAdmin.DELETE("/counselors/:id/schedules", adminScheduleHandler.Delete)
		restrictAdmin.PUT("/counselors/:id/schedules", adminScheduleHandler.Update)

		restrictAdmin.GET("/careers", adminCareerHandler.GetAll)
		restrictAdmin.POST("/careers", adminCareerHandler.Create)
		restrictAdmin.GET("/careers/:id", adminCareerHandler.GetById)
		restrictAdmin.PUT("/careers/:id", adminCareerHandler.Update)
		restrictAdmin.DELETE("/careers/:id", adminCareerHandler.Delete)

		restrictAdmin.GET("/users", adminUsersHandler.GetAll)
		restrictAdmin.GET("/users/:id", adminUsersHandler.GetById)
		restrictAdmin.DELETE("/users/:id", adminUsersHandler.Delete)

		restrictAdmin.DELETE("/forums/:id", forumAdminH.Delete)
	}

	// ssl
	// e.Logger.Fatal(e.StartTLS(":8080", "./ssl/certificate.crt", "./ssl/private.key"))

	e.Logger.Fatal(e.Start(":8080"))
}
