package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/matheusvidal21/smart-news-fetcher/internal/interfaces"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	articleService interfaces.ArticleServiceInterface
}

func NewArticleHandler(articleService interfaces.ArticleServiceInterface) *ArticleHandler {
	return &ArticleHandler{articleService: articleService}
}

func (h *ArticleHandler) FindAll(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	sort := c.DefaultQuery("sort", "asc")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
	}

	articles, err := h.articleService.FindAll(page, limit, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	logger.Info("Articles found: " + strconv.Itoa(len(articles)))
	c.JSON(http.StatusOK, articles)
}

func (h *ArticleHandler) FindOne(c *gin.Context) {
	id := c.Param("id")
	article, err := h.articleService.FindOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Article found: " + id)
	c.JSON(http.StatusOK, article)
}

func (h *ArticleHandler) Create(c *gin.Context) {
	var article dto.CreateArticleInput

	if err := c.BindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dto.Validate.Struct(article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdArticle, err := h.articleService.Create(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Article created: " + createdArticle.Title)
	c.JSON(http.StatusOK, createdArticle)
}

func (h *ArticleHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var article dto.UpdateArticleInput
	if err := c.BindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dto.Validate.Struct(article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uptdatedArticle, err := h.articleService.Update(id, article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Article updated: " + id)
	c.JSON(http.StatusOK, uptdatedArticle)
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.articleService.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Article deleted: " + id)
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}

func (h *ArticleHandler) FindBySourceID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid source ID"})
		return
	}

	articles, err := h.articleService.FindAllBySourceId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Articles found: " + strconv.Itoa(len(articles)))
	c.JSON(http.StatusOK, articles)
}
