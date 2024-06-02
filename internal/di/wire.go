//go:build wireinject
// +build wireinject

package di

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/matheusvidal21/smart-news-fetcher/internal/articles"
	"github.com/matheusvidal21/smart-news-fetcher/internal/auth"
	"github.com/matheusvidal21/smart-news-fetcher/internal/fetcher"
	"github.com/matheusvidal21/smart-news-fetcher/internal/sources"
	"github.com/matheusvidal21/smart-news-fetcher/internal/user"
)

var setFetcherDependecy = wire.NewSet(
	fetcher.NewFetcher,
	wire.Bind(new(fetcher.FetcherInterface), new(*fetcher.Fetcher)),
)

var setSourceHandlerDependecy = wire.NewSet(
	sources.NewSourceHandler,
	wire.Bind(new(sources.SourceHandlerInterface), new(*sources.SourceHandler)),
)

var setSourceServiceDependecy = wire.NewSet(
	sources.NewSourceService,
	wire.Bind(new(sources.SourceServiceInterface), new(*sources.SourceService)),
)

var setSourceRepositoryDependecy = wire.NewSet(
	sources.NewSourceRepository,
	wire.Bind(new(sources.SourceRepositoryInterface), new(*sources.SourceRepository)),
)

var setArticleHandlerDependecy = wire.NewSet(
	articles.NewArticleHandler,
	wire.Bind(new(articles.ArticleHandlerInterface), new(*articles.ArticleHandler)),
)

var setArticleServiceDependecy = wire.NewSet(
	articles.NewArticleService,
	wire.Bind(new(articles.ArticleServiceInterface), new(*articles.ArticleService)),
)

var setArticleRepositoryDependecy = wire.NewSet(
	articles.NewArticleRepository,
	wire.Bind(new(articles.ArticleRepositoryInterface), new(*articles.ArticleRepository)),
)

var setUserRepositoryDependecy = wire.NewSet(
	user.NewUserRepository,
	wire.Bind(new(user.UserRepositoryInterface), new(*user.UserRepository)),
)

var setUserServiceDependecy = wire.NewSet(
	user.NewUserService,
	wire.Bind(new(user.UserServiceInterface), new(*user.UserService)),
)

var setUserHandlerDependecy = wire.NewSet(
	user.NewUserHandler,
	wire.Bind(new(user.UserHandlerInterface), new(*user.UserHandler)),
)

func NewArticleHandler(db *sql.DB) *articles.ArticleHandler {
	wire.Build(
		setArticleRepositoryDependecy,
		setArticleServiceDependecy,
		setArticleHandlerDependecy,
	)
	return &articles.ArticleHandler{}
}

func NewSourceHandler(db *sql.DB) *sources.SourceHandler {
	wire.Build(
		setSourceRepositoryDependecy,
		setArticleRepositoryDependecy,
		setArticleServiceDependecy,
		setFetcherDependecy,
		setSourceServiceDependecy,
		setSourceHandlerDependecy,
	)

	return &sources.SourceHandler{}
}

func NewUserHandler(db *sql.DB, jwtService auth.JWTServiceInterface) *user.UserHandler {
	wire.Build(
		setUserRepositoryDependecy,
		setUserServiceDependecy,
		setUserHandlerDependecy,
	)
	return &user.UserHandler{}
}
