package articles

import (
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
)

type ArticleServiceInterface interface {
	FindAll(page, limit int, sort string) ([]Article, error)
	FindOne(id int) (dto.FindOneArticleOutput, error)
	Create(articleDto dto.CreateArticleInput) (dto.CreateArticleOutput, error)
	Update(id int, articleDto dto.UpdateArticleInput) (dto.UpdateArticleOutput, error)
	Delete(id int) error
	FindAllBySourceId(sourceID int) ([]Article, error)
}

type ArticleService struct {
	articleRepository ArticleRepositoryInterface
}

func NewArticleService(articleRepository ArticleRepositoryInterface) *ArticleService {
	return &ArticleService{articleRepository: articleRepository}
}

func (as *ArticleService) FindAll(page, limit int, sort string) ([]Article, error) {
	return as.articleRepository.FindAll(page, limit, sort)
}

func (as *ArticleService) FindOne(id int) (dto.FindOneArticleOutput, error) {
	article, err := as.articleRepository.FindOne(id)

	if err != nil {
		return dto.FindOneArticleOutput{}, err
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
	article := Article{
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
		return dto.CreateArticleOutput{}, err
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
	article := Article{
		Title:       articleDto.Title,
		Description: articleDto.Description,
		Content:     articleDto.Content,
		Link:        articleDto.Link,
		PubDate:     articleDto.PubDate,
		SourceID:    articleDto.SourceID,
	}

	articleSaved, err := as.articleRepository.Update(id, article)

	if err != nil {
		return dto.UpdateArticleOutput{}, err
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
	return as.articleRepository.Delete(id)
}

func (as *ArticleService) FindAllBySourceId(sourceID int) ([]Article, error) {
	return as.articleRepository.FindAllBySourceId(sourceID)
}
