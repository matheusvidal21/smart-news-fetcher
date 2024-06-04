package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/matheusvidal21/smart-news-fetcher/internal/models"
)

type SourceRepositoryInterface interface {
	FindAll(page, limit int, sort string) ([]models.Source, error)
	FindOne(id int) (models.Source, error)
	Create(source models.Source) (models.Source, error)
	Update(id int, source models.Source) (models.Source, error)
	Delete(id int) error
	FindByUrl(url string) (models.Source, error)
	FindByUserId(userId int) ([]models.Source, error)
	FindAllActive() ([]models.Source, error)
}

type SourceServiceInterface interface {
	FindAll(page, limit int, sort string) ([]models.Source, error)
	FindOne(id int) (dto.FindOneSourceOutput, error)
	Create(sourceDto dto.CreateSourceInput) (dto.CreateSourceOutput, error)
	Update(id int, sourceDto dto.UpdateSourceInput) (dto.UpdateSourceOutput, error)
	Delete(id int) error
	LoadFeed(id int) error
	FindByUserId(userId int) ([]models.Source, error)
	SubscribeToNewsletter(id int) error
	UnsubscribeFromNewsletter(id int) error
	StartSubscription(source models.Source)
	InitializeSubscription()
}

type SourceHandlerInterface interface {
	FindAll(c *gin.Context)
	FindOne(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	LoadFeed(c *gin.Context)
	FindByUserId(c *gin.Context)
	SubscribeToNewsletter(c *gin.Context)
	UnsubscribeFromNewsletter(c *gin.Context)
}
