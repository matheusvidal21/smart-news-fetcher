package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/matheusvidal21/smart-news-fetcher/internal/interfaces"
	"net/http"
	"strconv"
)

type SourceHandler struct {
	sourceService interfaces.SourceServiceInterface
}

func NewSourceHandler(sourceService interfaces.SourceServiceInterface) *SourceHandler {
	return &SourceHandler{sourceService: sourceService}
}

func (sh *SourceHandler) FindAll(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	sort := c.DefaultQuery("sort", "asc")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	sources, err := sh.sourceService.FindAll(page, limit, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Sources found: " + strconv.Itoa(len(sources)))
	c.JSON(http.StatusOK, sources)
}

func (sh *SourceHandler) FindOne(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid source ID"})
		return
	}

	source, err := sh.sourceService.FindOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Source found: " + idStr)
	c.JSON(http.StatusOK, source)
}

func (sh *SourceHandler) Create(c *gin.Context) {
	var source dto.CreateSourceInput
	if err := c.BindJSON(&source); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dto.Validate.Struct(source); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdSource, err := sh.sourceService.Create(source)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Source created: " + createdSource.Name)
	c.JSON(http.StatusOK, createdSource)
}

func (sh *SourceHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid source ID"})
		return
	}

	var source dto.UpdateSourceInput
	if err = c.BindJSON(&source); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dto.Validate.Struct(source); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uptadedSource, err := sh.sourceService.Update(id, source)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Source updated: " + idStr)
	c.JSON(http.StatusOK, uptadedSource)
}

func (sh *SourceHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid source ID"})
		return
	}

	err = sh.sourceService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Source deleted: " + idStr)
	c.JSON(http.StatusOK, gin.H{"message": "Source deleted successfully"})
}

func (sh *SourceHandler) LoadFeed(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid source ID"})
		return
	}

	err = sh.sourceService.LoadFeed(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Feed loaded: " + idStr)
	c.JSON(http.StatusOK, gin.H{"message": "Feed loaded successfully"})
}

func (sh *SourceHandler) FindByUserId(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	sources, err := sh.sourceService.FindByUserId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Sources found: " + strconv.Itoa(len(sources)))
	c.JSON(http.StatusOK, sources)
}
