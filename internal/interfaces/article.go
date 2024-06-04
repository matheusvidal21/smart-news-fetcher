package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/matheusvidal21/smart-news-fetcher/internal/models"
)

type ArticleRepositoryInterface interface {
	FindAll(page, limit int, sort string) ([]models.Article, error)
	FindOne(id string) (models.Article, error)
	Create(article models.Article) (models.Article, error)
	Update(id string, article models.Article) (models.Article, error)
	Delete(id string) error
	FindAllBySourceId(sourceID int) ([]models.Article, error)
}

type ArticleServiceInterface interface {
	GenerateArticleID(title, link string) string
	FindAll(page, limit int, sort string) ([]models.Article, error)
	FindOne(id string) (dto.FindOneArticleOutput, error)
	Create(articleDto dto.CreateArticleInput) (dto.CreateArticleOutput, error)
	Update(id string, articleDto dto.UpdateArticleInput) (dto.UpdateArticleOutput, error)
	Delete(id string) error
	FindAllBySourceId(sourceID int) ([]models.Article, error)
}

type ArticleHandlerInterface interface {
	FindAll(c *gin.Context)
	FindOne(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindBySourceID(c *gin.Context)
}
