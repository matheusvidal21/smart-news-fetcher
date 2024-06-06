package main

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	configs "github.com/matheusvidal21/smart-news-fetcher/configs"
	"github.com/matheusvidal21/smart-news-fetcher/internal/auth"
	"github.com/matheusvidal21/smart-news-fetcher/internal/di"
	"github.com/matheusvidal21/smart-news-fetcher/internal/email"
	"github.com/matheusvidal21/smart-news-fetcher/internal/interfaces"
	"github.com/matheusvidal21/smart-news-fetcher/internal/middleware"
	"strconv"
)

type Server struct {
	Config        *configs.Conf
	DB            *sql.DB
	JWTService    auth.JWTServiceInterface
	EmailService  interfaces.EmailService
	SourceService interfaces.SourceServiceInterface
	Router        *gin.Engine
}

func NewServer() (*Server, error) {
	conf := configs.LoadConfigs(".")
	if conf == nil {
		return nil, errors.New("failed to load configs")
	}

	db, err := initDB(conf.DBDriver, conf.DBSource)
	if err != nil {
		return nil, errors.New("failed to connect to database: " + err.Error())
	}

	jwtExpiration, _ := strconv.Atoi(conf.JWTExpirationMinutes)
	smtpPort, _ := strconv.Atoi(conf.SMTP_PORT)
	jwtService := auth.NewJWTService(conf.JWTSecretKey, jwtExpiration)
	emailService := email.NewEmailService(conf.SMTP_HOST, smtpPort, conf.SMTP_USER, conf.SMTP_PASSWORD, conf.SMTP_FROM_EMAIL)
	sourceService := di.NewSourceService(db, jwtService, emailService)

	router := gin.Default()
	server := &Server{
		Config:        conf,
		DB:            db,
		JWTService:    jwtService,
		EmailService:  emailService,
		SourceService: sourceService,
		Router:        router,
	}
	return server, nil
}

func (s *Server) InitializeRoutes() {
	articleHandler := di.NewArticleHandler(s.DB)
	sourceHandler := di.NewSourceHandler(s.DB, s.JWTService, s.EmailService)
	userHandler := di.NewUserHandler(s.DB, s.JWTService, s.EmailService)

	s.SourceService.InitializeSubscription()

	articles := s.Router.Group("/articles")
	{
		articles.Use(middleware.AuthMiddleware(s.JWTService))
		articles.GET("/", articleHandler.FindAll)
		articles.GET("/:id", articleHandler.FindOne)
		articles.POST("/", articleHandler.Create)
		articles.PUT("/:id", articleHandler.Update)
		articles.DELETE("/:id", articleHandler.Delete)
		articles.GET("/source/:id", articleHandler.FindBySourceID)
	}

	sources := s.Router.Group("/sources")
	{
		sources.Use(middleware.AuthMiddleware(s.JWTService))
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

	users := s.Router.Group("/users")
	{
		users.POST("/", userHandler.CreateUser)
		users.GET("/find_by_email/:email", userHandler.FindByEmail)
		users.GET("/:id", userHandler.FindById)
		users.DELETE("/:email", userHandler.DeleteUser)
		users.POST("/login", userHandler.Login)
		users.POST("/update_password", userHandler.UpdatePassword)
	}
}

func (s *Server) Run() error {
	return s.Router.Run(s.Config.WebServerPort)
}

func initDB(driver, source string) (*sql.DB, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
