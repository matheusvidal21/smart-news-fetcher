package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/matheusvidal21/smart-news-fetcher/internal/models"
)

type ArticleRepositoryInterface interface {
	FindAll(page, limit int, sort string) ([]models.Article, error)
	FindOne(id int) (models.Article, error)
	Create(article models.Article) (models.Article, error)
	Update(id int, article models.Article) (models.Article, error)
	Delete(id int) error
	FindAllBySourceId(sourceID int) ([]models.Article, error)
}

type ArticleServiceInterface interface {
	FindAll(page, limit int, sort string) ([]models.Article, error)
	FindOne(id int) (dto.FindOneArticleOutput, error)
	Create(articleDto dto.CreateArticleInput) (dto.CreateArticleOutput, error)
	Update(id int, articleDto dto.UpdateArticleInput) (dto.UpdateArticleOutput, error)
	Delete(id int) error
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
