package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	configs "github.com/matheusvidal21/smart-news-fetcher/configs"
	"github.com/matheusvidal21/smart-news-fetcher/internal/auth"
	"github.com/matheusvidal21/smart-news-fetcher/internal/di"
	"github.com/matheusvidal21/smart-news-fetcher/internal/email"
	"github.com/matheusvidal21/smart-news-fetcher/internal/middleware"
	"github.com/matheusvidal21/smart-news-fetcher/pkg/logger"
	"log"
	"strconv"
)

func main() {
	if err := logger.InitializeLogger(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.CloseLogger()

	conf := configs.LoadConfigs(".")

	db, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	jwtExpiration, _ := strconv.Atoi(conf.JWTExpirationMinutes)
	smtpPort, _ := strconv.Atoi(conf.SMTP_PORT)
	jwtService := auth.NewJWTService(conf.JWTSecretKey, jwtExpiration)
	emailService := email.NewEmailService(conf.SMTP_HOST, smtpPort, conf.SMTP_USER, conf.SMTP_PASSWORD, conf.SMTP_FROM_EMAIL)
	sourceService := di.NewSourceService(db, jwtService, emailService)

	articleHandler := di.NewArticleHandler(db)
	sourceHandler := di.NewSourceHandler(db, jwtService, emailService)
	userHandler := di.NewUserHandler(db, jwtService, emailService)
	router := gin.Default()
	sourceService.InitializeSubscription()

	articles := router.Group("/articles")
	{
		articles.Use(middleware.AuthMiddleware(jwtService))
		articles.GET("/", articleHandler.FindAll)
		articles.GET("/:id", articleHandler.FindOne)
		articles.POST("/", articleHandler.Create)
		articles.PUT("/:id", articleHandler.Update)
		articles.DELETE("/:id", articleHandler.Delete)
		articles.GET("/source/:id", articleHandler.FindBySourceID)
	}

	sources := router.Group("/sources")
	{
		sources.Use(middleware.AuthMiddleware(jwtService))
		sources.GET("/", sourceHandler.FindAll)
		sources.GET("/:id", sourceHandler.FindOne)
		sources.POST("/", sourceHandler.Create)
		sources.PUT("/:id", sourceHandler.Update)
		sources.DELETE("/:id", sourceHandler.Delete)
		sources.GET("/load_feed/:id", sourceHandler.LoadFeed)
		sources.GET("/find_by_user/:id", sourceHandler.FindByUserId)
		sources.GET("/subscribe/:id", sourceHandler.SubscribeToNewsletter)
		sources.GET("/unsubscribe/:id", sourceHandler.UnsubscribeFromNewsletter)
	}

	users := router.Group("/users")
	{
		users.POST("/", userHandler.CreateUser)
		users.GET("/find_by_email/:email", userHandler.FindByEmail)
		users.GET("/:id", userHandler.FindById)
		users.DELETE("/:email", userHandler.DeleteUser)
		users.POST("/login", userHandler.Login)
		users.POST("/update_password", userHandler.UpdatePassword)
	}

	router.Run(conf.WebServerPort)
}
