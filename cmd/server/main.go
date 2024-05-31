package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	configs "github.com/matheusvidal21/smart-news-fetcher/configs"
	"github.com/matheusvidal21/smart-news-fetcher/internal/di"
	"github.com/matheusvidal21/smart-news-fetcher/pkg/logger"
	"log"
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

	articleHandler := di.NewArticleHandler(db)
	sourceHandler := di.NewSourceHandler(db)
	router := gin.Default()

	//fetcher := di.NewFetcher(db)
	//	fetcher.StartScheduler(5 * time.Second)

	articles := router.Group("/articles")
	{
		articles.GET("/", articleHandler.FindAll)
		articles.GET("/:id", articleHandler.FindOne)
		articles.POST("/", articleHandler.Create)
		articles.PUT("/:id", articleHandler.Update)
		articles.DELETE("/:id", articleHandler.Delete)
		articles.GET("/source/:id", articleHandler.FindBySourceID)
	}

	sources := router.Group("/sources")
	{
		sources.GET("/", sourceHandler.FindAll)
		sources.GET("/:id", sourceHandler.FindOne)
		sources.POST("/", sourceHandler.Create)
		sources.PUT("/:id", sourceHandler.Update)
		sources.DELETE("/:id", sourceHandler.Delete)
	}

	router.Run(conf.WebServerPort)
}
