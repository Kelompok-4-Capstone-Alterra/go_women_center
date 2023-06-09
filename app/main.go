package main

import (
	"log"
	"os"

	AdminAuthHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/handler"
	adminAuthMidd "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/handler/middleware"
	AdminAuthRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/repository"
	AdminAuthUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/auth/usecase"

	CounselorAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/handler"
	CounselorAdminRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/repository"
	CounselorAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/usecase"

	AdminScheduleHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/schedule/handler"
	AdminScheduleRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/schedule/repository"
	AdminScheduleUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/schedule/usecase"

	AdminStatisticHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/statistics/handler"
	AdminStatisticRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/statistics/repository"
	AdminStatisticUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/statistics/usecase"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/app/config"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	TopicHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/topic/handler"

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

	UserScheduleHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/schedule/handler"
	UserScheduleRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/schedule/repository"
	UserScheduleUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/schedule/usecase"

	ForumUserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/handler"
	ForumUserRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/repository"
	ForumUserUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/forum/usecase"
	UserForumAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/handler"
	UserForumAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/repository"
	UserForumAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/user_forum/usecase"

	ForumAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum/handler"
	ForumAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum/repository"
	ForumAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/forum/usecase"

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

	VoucherUserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/voucher/handler"
	VoucherUserRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/voucher/repository"
	VoucherUserUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/voucher/usecase"

	TransactionUserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction/handler"
	TransactionUserRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction/repository"
	TransactionUserUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction/usecase"

	TransactionAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction/handler"
	TransactionAdminRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction/repository"
	TransactionAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction/usecase"

	ArticleAdminHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/article/handler"
	ArticleAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/article/repository"
	ArticleAdminUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/article/usecase"
	CommentAdminRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/comment/repository"

	ArticleUserHandler "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/article/handler"
	ArticleUserRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/article/repository"
	ArticleUserUsecase "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/article/usecase"
	CommentUserRepository "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/comment/repository"

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

	mailConf := helper.NewEmailSender(
		587,
		"smtp.gmail.com",
		os.Getenv("CONFIG_AUTH_EMAIL"),
		os.Getenv("CONFIG_AUTH_PASSWORD"),
		"Women Center <averilprimayuda@gmail.com>", //TODO: set email to the proper one
	)

	db := dbconf.InitDB()
	sslconf.InitSSL()

	// helper
	jwtConf := helper.NewAuthJWT(os.Getenv("JWT_SECRET_USER"), os.Getenv("JWT_SECRET_ADMIN"))
	encryptor := helper.NewEncryptor()
	otpGenerator := helper.NewOtpGenerator()
	googleUUID := helper.NewGoogleUUID()
	image := helper.NewImage("women-center", googleUUID)
	log.Print(db, googleUUID)

	// handler
	userAuthRepo := UserAuthRepo.NewUserRepository(db)
	otpRepo := UserAuthRepo.NewLocalCache(config.CLEANUP_INTERVAL)
	userAuthUsecase := UserAuthUsecase.NewUserUsecase(userAuthRepo, googleUUID, &mailConf, otpRepo, otpGenerator, encryptor)
	userAuthHandler := UserAuthHandler.NewUserHandler(userAuthUsecase, googleOauthConfig, jwtConf)

	userRepo := UserProfileRepo.NewMysqlUserRepository(db)
	userUsecase := UserProfileUsecase.NewProfileUsecase(userRepo, image, encryptor)
	userHandler := UserProfileHandler.NewProfileHandler(userUsecase)

	userCareerRepo := CareerUserRepository.NewMysqlCareerRepository(db)
	userCareerUsecase := CareerUserUsecase.NewCareerUsecase(userCareerRepo)
	userCareerHandler := CareerUserHandler.NewCareerHandler(userCareerUsecase)

	userArticleRepo := ArticleUserRepository.NewMysqlArticleRepository(db)
	userCommentRepo := CommentUserRepository.NewMysqlArticleRepository(db)
	userArticleUsecase := ArticleUserUsecase.NewArticleUsecase(userArticleRepo, userCommentRepo, userAuthRepo, googleUUID)
	userArticleHandler := ArticleUserHandler.NewArticleHandler(userArticleUsecase)

	adminAuthRepo := AdminAuthRepo.NewAdminRepo(db)
	adminAuthUsecase := AdminAuthUsecase.NewAuthUsecase(adminAuthRepo, encryptor)
	adminAuthHandler := AdminAuthHandler.NewAuthHandler(adminAuthUsecase, jwtConf)

	adminCareerRepo := CareerAdminRepository.NewMysqlCareerRepository(db)
	adminCareerUsecase := CareerAdminUsecase.NewCareerUsecase(adminCareerRepo, image, googleUUID)
	adminCareerHandler := CareerAdminHandler.NewCareerHandler(adminCareerUsecase)

	adminUsersRepo := UsersAdminRepository.NewMysqlUserRepository(db)
	adminUsersUsecase := UsersAdminUsecase.NewUserUsecase(adminUsersRepo, image)
	adminUsersHandler := UsersAdminHandler.NewUserHandler(adminUsersUsecase)

	adminCounselorRepo := CounselorAdminRepo.NewMysqlCounselorRepository(db)
	adminScheduleRepo := AdminScheduleRepo.NewMysqlScheduleRepository(db)
	adminCounselorUsecase := CounselorAdminUsecase.NewCounselorUsecase(adminCounselorRepo, adminScheduleRepo, image, googleUUID)
	adminCounselorHandler := CounselorAdminHandler.NewCounselorHandler(adminCounselorUsecase)
	adminScheduleUsecase := AdminScheduleUsecase.NewScheduleUsecase(adminCounselorRepo, adminScheduleRepo, googleUUID)
	adminScheduleHandler := AdminScheduleHandler.NewScheduleHandler(adminScheduleUsecase)

	userForumR := UserForumAdminRepository.NewMysqlUserForumRepository(db)
	userForumU := UserForumAdminUsecase.NewUserForumUsecase(userForumR)
	userForumH := UserForumAdminHandler.NewUserForumHandler(userForumU)

	forumR := ForumUserRepository.NewMysqlForumRepository(db)
	forumU := ForumUserUsecase.NewForumUsecase(forumR, userForumR)
	forumH := ForumUserHandler.NewForumHandler(forumU)

	forumAdminR := ForumAdminRepository.NewMysqlForumAdminRepository(db)
	forumAdminU := ForumAdminUsecase.NewForumAdminUsecase(forumAdminR)
	forumAdminH := ForumAdminHandler.NewForumAdminHandler(forumAdminU)

	adminCommentRepo := CommentAdminRepository.NewMysqlArticleRepository(db)
	adminArticleRepo := ArticleAdminRepository.NewMysqlArticleRepository(db)
	adminArticleUsecase := ArticleAdminUsecase.NewArticleUsecase(adminArticleRepo, adminCommentRepo, userAuthRepo, image, googleUUID)
	adminArticleHandler := ArticleAdminHandler.NewArticleHandler(adminArticleUsecase)

	ReadingListR := ReadingListRepository.NewMysqlReadingListRepository(db)
	ReadingListU := ReadingListUsecase.NewReadingListUsecase(ReadingListR)
	ReadingListH := ReadingListHandler.NewReadingListHandler(ReadingListU)

	ReadingListArticleR := ReadingListArticleRepository.NewMysqlReadingListArticleRepository(db)
	ReadingListArticleU := ReadingListArticleUsecase.NewReadingListArticleUsecase(ReadingListArticleR, ReadingListR)
	ReadingListArticleH := ReadingListArticleHandler.NewReadingListArticleHandler(ReadingListArticleU)

	topicHandler := TopicHandler.NewTopicHandler()

	midtransServerKey := os.Getenv("MIDTRANS_SERVER_KEY")
	midtransNotifHandler := os.Getenv("MIDTRANS_NOTIFICATION_HANDLER")
	log.Println("====MIDTRANS NOTIF HANDLER =", midtransNotifHandler, "====")

	userVoucherRepo := VoucherUserRepo.NewMysqltransactionRepository(db)
	userVoucherUsecase := VoucherUserUsecase.NewtransactionUsecase(userVoucherRepo)
	userVoucherHandler := VoucherUserHandler.NewVoucherHandler(userVoucherUsecase)

	userCounselorRepo := CounselorUserRepo.NewMysqlCounselorRepository(db)
	userScheduleRepo := UserScheduleRepo.NewMysqlScheduleRepository(db)

	userTransactionRepo := TransactionUserRepo.NewMysqltransactionRepository(db)
	userTransactionUsecase := TransactionUserUsecase.NewtransactionUsecase(midtransServerKey, googleUUID, userTransactionRepo, midtransNotifHandler, userCounselorRepo, userScheduleRepo, userVoucherRepo)
	userTransactionHandler := TransactionUserHandler.NewTransactionHandler(userTransactionUsecase)

	userScheduleUseCase := UserScheduleUsecase.NewScheduleUsecase(userScheduleRepo, userTransactionRepo, userCounselorRepo)
	userScheduleHandler := UserScheduleHandler.NewScheduleHandler(userScheduleUseCase)

	userReviewRepo := CounselorUserRepo.NewMysqlReviewRepository(db)
	userCounselorUsecase := CounselorUserUsecase.NewCounselorUsecase(userCounselorRepo, userReviewRepo, userAuthRepo, userTransactionRepo, googleUUID)
	userCounselorHandler := CounselorUserHandler.NewCounselorHandler(userCounselorUsecase)

	adminTransactionRepo := TransactionAdminRepo.NewMysqltransactionRepository(db)
	adminVoucherRepo := TransactionAdminRepo.NewMysqlVoucherRepository(db)
	adminTransactionUsecase := TransactionAdminUsecase.NewtransactionUsecase(adminTransactionRepo, adminVoucherRepo, googleUUID)
	adminTransactionHandler := TransactionAdminHandler.NewTransactionHandler(adminTransactionUsecase)

	adminStatisticRepo := AdminStatisticRepo.NewStatisticGormMysqlRepo(db)
	adminStatisticUsecase := AdminStatisticUsecase.NewAuthUsecase(adminStatisticRepo)
	adminStatisticHandler := AdminStatisticHandler.NewStatisticController(adminStatisticUsecase)

	e := echo.New()

	// middleware echo
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	// payment callback
	midtrans := e.Group("/midtrans")
	{
		midtrans.POST("/transaction/callback", userTransactionHandler.MidtransNotification)
	}

	google := e.Group("/google")
	{
		google.GET("/google/login", userAuthHandler.LoginGoogleHandler)
		google.GET("/google/callback", userAuthHandler.LoginGoogleCallback)
	}

	users := e.Group("/users")
	{
		public := users.Group("/public")
		{
			public.GET("/counselors", userCounselorHandler.GetAll)
			public.GET("/careers", userCareerHandler.GetAll)
			public.GET("/careers/:id", userCareerHandler.GetById)
			public.GET("/articles", userArticleHandler.GetAll)
			public.GET("/articles/:id", userArticleHandler.GetById)
			public.GET("/topics", topicHandler.GetAll)
			public.GET("/forums", forumH.GetAll)
			public.GET("/forums/:id", forumH.GetById)
		}

		auth := users.Group("/auth")

		{
			auth.POST("/verify/unique", userAuthHandler.VerifyUniqueCredential)
			auth.POST("/verify", userAuthHandler.VerifyEmailHandler)
			auth.POST("/register", userAuthHandler.RegisterHandler)
			auth.POST("/login", userAuthHandler.LoginHandler)
			auth.POST("/verify/forget", userAuthHandler.CheckIsRegistered)
			auth.POST("/forget-password", userAuthHandler.ForgetPassword)
		}

		restrictUsers := users.Group("", userAuthMidd.JWTUser(), userAuthMidd.CheckUser(userAuthUsecase))

		{
			restrictUsers.GET("/profile", userHandler.GetById)
			restrictUsers.PUT("/profile", userHandler.Update)
			restrictUsers.PUT("/profile/password", userHandler.UpdatePassword)

			restrictUsers.GET("/counselors/:id", userCounselorHandler.GetById)
			restrictUsers.POST("/counselors/:id/reviews", userCounselorHandler.CreateReview)
			restrictUsers.GET("/counselors/:id/reviews", userCounselorHandler.GetAllReview)
			restrictUsers.GET("/counselors/:id/schedules", userScheduleHandler.GetCurrSchedule)

			restrictUsers.GET("/forums", forumH.GetAll)
			restrictUsers.GET("/forums/:id", forumH.GetById)
			restrictUsers.POST("/forums", forumH.Create)
			restrictUsers.PUT("/forums/:id", forumH.Update)
			restrictUsers.DELETE("/forums/:id", forumH.Delete)
			restrictUsers.POST("/forums/joins", userForumH.Create)

			restrictUsers.GET("/reading-lists", ReadingListH.GetAll)
			restrictUsers.GET("/reading-lists/:id", ReadingListH.GetById)
			restrictUsers.POST("/reading-lists", ReadingListH.Create)
			restrictUsers.PUT("/reading-lists/:id", ReadingListH.Update)
			restrictUsers.DELETE("/reading-lists/:id", ReadingListH.Delete)

			restrictUsers.POST("/reading-lists/save", ReadingListArticleH.Create)
			restrictUsers.DELETE("/reading-lists/save/:id", ReadingListArticleH.Delete)

			restrictUsers.GET("/articles", userArticleHandler.GetAll)
			restrictUsers.GET("/articles/:id", userArticleHandler.GetById)
			restrictUsers.POST("/articles/:id/comments", userArticleHandler.CreateComment)
			restrictUsers.GET("/articles/:id/comments", userArticleHandler.GetAllComment)
			restrictUsers.DELETE("/articles/:article_id/comments/:comment_id", userArticleHandler.DeleteComment)

			restrictUsers.GET("/vouchers", userVoucherHandler.GetAll)
			restrictUsers.GET("/transactions", userTransactionHandler.GetAllTransaction)
			restrictUsers.POST("/transactions", userTransactionHandler.SendTransaction)
			restrictUsers.GET("/transactions/:id", userTransactionHandler.GetTransactionDetail)
			restrictUsers.POST("/transactions/join", userTransactionHandler.UserJoinHandler)
		}

	}

	admin := e.Group("/admin")
	{
		auth := admin.Group("/auth")
		{
			auth.POST("/login", adminAuthHandler.LoginHandler)
		}

		restrictAdmin := admin.Group("", adminAuthMidd.JWTAdmin())
		{
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

			restrictAdmin.GET("/articles", adminArticleHandler.GetAll)
			restrictAdmin.POST("/articles", adminArticleHandler.Create)
			restrictAdmin.GET("/articles/:id", adminArticleHandler.GetById)
			restrictAdmin.PUT("/articles/:id", adminArticleHandler.Update)
			restrictAdmin.DELETE("/articles/:id", adminArticleHandler.Delete)
			restrictAdmin.GET("/articles/:id/comments", adminArticleHandler.GetAllComment)
			restrictAdmin.DELETE("/articles/:article_id/comments/:comment_id", adminArticleHandler.DeleteComment)

			restrictAdmin.GET("/users", adminUsersHandler.GetAll)
			restrictAdmin.GET("/users/:id", adminUsersHandler.GetById)
			restrictAdmin.DELETE("/users/:id", adminUsersHandler.Delete)

			restrictAdmin.GET("/forums", forumAdminH.GetAll)
			restrictAdmin.GET("/forums/:id", forumAdminH.GetById)
			restrictAdmin.DELETE("/forums/:id", forumAdminH.Delete)

			restrictAdmin.GET("/transactions", adminTransactionHandler.GetAll)
			restrictAdmin.PUT("/transactions/link", adminTransactionHandler.SendLink)
			restrictAdmin.PUT("/transactions/cancel", adminTransactionHandler.CancelTransaction)

			restrictAdmin.GET("/transactions/report", adminTransactionHandler.GetReport)
			restrictAdmin.GET("/transactions/report/download", adminTransactionHandler.DownloadReport)

			restrictAdmin.GET("/statistics", adminStatisticHandler.GetData)
		}

	}

	// ssl
	e.Logger.Fatal(e.StartTLS(":8080", "./ssl/certificate.crt", "./ssl/private.key"))

	// e.Logger.Fatal(e.Start(":8080"))

}
