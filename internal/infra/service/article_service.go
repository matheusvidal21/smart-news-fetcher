package service

import (
	"errors"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/matheusvidal21/smart-news-fetcher/internal/interfaces"
	"github.com/matheusvidal21/smart-news-fetcher/internal/models"
)

type ArticleService struct {
	articleRepository interfaces.ArticleRepositoryInterface
}

func NewArticleService(articleRepository interfaces.ArticleRepositoryInterface) *ArticleService {
	return &ArticleService{articleRepository: articleRepository}
}

func (as *ArticleService) FindAll(page, limit int, sort string) ([]models.Article, error) {
	articles, err := as.articleRepository.FindAll(page, limit, sort)
	if err != nil {
		return nil, errors.New("Articles not found: " + err.Error())
	}
	return articles, nil
}

func (as *ArticleService) FindOne(id int) (dto.FindOneArticleOutput, error) {
	article, err := as.articleRepository.FindOne(id)

	if err != nil {
		return dto.FindOneArticleOutput{}, errors.New("Article not found: " + err.Error())
	}

	return dto.FindOneArticleOutput{
		ID:          article.ID,
		Title:       article.Title,
		Description: article.Description,
		Content:     article.Content,
		Link:        article.Link,
		PubDate:     article.PubDate,
		Author:      article.Author,
		SourceID:    article.SourceID,
	}, nil
}

func (as *ArticleService) Create(articleDto dto.CreateArticleInput) (dto.CreateArticleOutput, error) {
	article := models.Article{
		Title:       articleDto.Title,
		Description: articleDto.Description,
		Content:     articleDto.Content,
		Link:        articleDto.Link,
		PubDate:     articleDto.PubDate,
		Author:      articleDto.Author,
		SourceID:    articleDto.SourceID,
	}

	articleSaved, err := as.articleRepository.Create(article)

	if err != nil {
		return dto.CreateArticleOutput{}, errors.New("Article not created: " + err.Error())
	}

	return dto.CreateArticleOutput{
		ID:          articleSaved.ID,
		Title:       articleSaved.Title,
		Description: articleSaved.Description,
		Content:     articleSaved.Content,
		Link:        articleSaved.Link,
		PubDate:     articleSaved.PubDate,
		Author:      articleSaved.Author,
		SourceID:    articleSaved.SourceID,
	}, nil
}

func (as *ArticleService) Update(id int, articleDto dto.UpdateArticleInput) (dto.UpdateArticleOutput, error) {
	article := models.Article{
		Title:       articleDto.Title,
		Description: articleDto.Description,
		Content:     articleDto.Content,
		Link:        articleDto.Link,
		PubDate:     articleDto.PubDate,
		SourceID:    articleDto.SourceID,
	}

	articleSaved, err := as.articleRepository.Update(id, article)

	if err != nil {
		return dto.UpdateArticleOutput{}, errors.New("Article not updated: " + err.Error())
	}

	return dto.UpdateArticleOutput{
		ID:          articleSaved.ID,
		Title:       articleSaved.Title,
		Description: articleSaved.Description,
		Content:     articleSaved.Content,
		Link:        articleSaved.Link,
		PubDate:     articleSaved.PubDate,
		Author:      articleSaved.Author,
		SourceID:    articleSaved.SourceID,
	}, nil
}

func (as *ArticleService) Delete(id int) error {
	err := as.articleRepository.Delete(id)
	if err != nil {
		return errors.New("Article not deleted: " + err.Error())
	}
	return nil
}

func (as *ArticleService) FindAllBySourceId(sourceID int) ([]models.Article, error) {
	articles, err := as.articleRepository.FindAllBySourceId(sourceID)

	if err != nil {
		return nil, errors.New("Articles not found: " + err.Error())
	}
	return articles, nil
}
