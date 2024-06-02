// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/matheusvidal21/smart-news-fetcher/internal/auth"
	"github.com/matheusvidal21/smart-news-fetcher/internal/fetcher"
	"github.com/matheusvidal21/smart-news-fetcher/internal/infra/database"
	"github.com/matheusvidal21/smart-news-fetcher/internal/infra/handler"
	"github.com/matheusvidal21/smart-news-fetcher/internal/infra/service"
	"github.com/matheusvidal21/smart-news-fetcher/internal/interfaces"
)

// Injectors from wire.go:

func NewArticleHandler(db *sql.DB) *handler.ArticleHandler {
	articleRepository := database.NewArticleRepository(db)
	articleService := service.NewArticleService(articleRepository)
	articleHandler := handler.NewArticleHandler(articleService)
	return articleHandler
}

func NewSourceHandler(db *sql.DB, jwtService auth.JWTServiceInterface, emailService interfaces.EmailService) *handler.SourceHandler {
	sourceRepository := database.NewSourceRepository(db)
	userRepository := database.NewUserRepository(db)
	userService := service.NewUserService(userRepository, jwtService)
	articleRepository := database.NewArticleRepository(db)
	articleService := service.NewArticleService(articleRepository)
	fetcherFetcher := fetcher.NewFetcher(articleService)
	sourceService := service.NewSourceService(sourceRepository, userService, emailService, fetcherFetcher)
	sourceHandler := handler.NewSourceHandler(sourceService)
	return sourceHandler
}

func NewUserHandler(db *sql.DB, jwtService auth.JWTServiceInterface) *handler.UserHandler {
	userRepository := database.NewUserRepository(db)
	userService := service.NewUserService(userRepository, jwtService)
	userHandler := handler.NewUserHandler(userService)
	return userHandler
}

// wire.go:

var setFetcherDependecy = wire.NewSet(fetcher.NewFetcher, wire.Bind(new(interfaces.FetcherInterface), new(*fetcher.Fetcher)))

var setSourceHandlerDependecy = wire.NewSet(handler.NewSourceHandler, wire.Bind(new(interfaces.SourceHandlerInterface), new(*handler.SourceHandler)))

var setSourceServiceDependecy = wire.NewSet(service.NewSourceService, wire.Bind(new(interfaces.SourceServiceInterface), new(*service.SourceService)))

var setSourceRepositoryDependecy = wire.NewSet(database.NewSourceRepository, wire.Bind(new(interfaces.SourceRepositoryInterface), new(*database.SourceRepository)))

var setArticleHandlerDependecy = wire.NewSet(handler.NewArticleHandler, wire.Bind(new(interfaces.ArticleHandlerInterface), new(*handler.ArticleHandler)))

var setArticleServiceDependecy = wire.NewSet(service.NewArticleService, wire.Bind(new(interfaces.ArticleServiceInterface), new(*service.ArticleService)))

var setArticleRepositoryDependecy = wire.NewSet(database.NewArticleRepository, wire.Bind(new(interfaces.ArticleRepositoryInterface), new(*database.ArticleRepository)))

var setUserRepositoryDependecy = wire.NewSet(database.NewUserRepository, wire.Bind(new(interfaces.UserRepositoryInterface), new(*database.UserRepository)))

var setUserServiceDependecy = wire.NewSet(service.NewUserService, wire.Bind(new(interfaces.UserServiceInterface), new(*service.UserService)))

var setUserHandlerDependecy = wire.NewSet(handler.NewUserHandler, wire.Bind(new(interfaces.UserHandlerInterface), new(*handler.UserHandler)))
