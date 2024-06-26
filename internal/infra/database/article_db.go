package database

import (
	"database/sql"
	"github.com/matheusvidal21/smart-news-fetcher/internal/models"
	"github.com/matheusvidal21/smart-news-fetcher/pkg/utils"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (ar *ArticleRepository) FindAll(page, limit int, sort string) ([]models.Article, error) {
	sql := "SELECT id, title, description, content, link, pub_date, author, source_id FROM articles"

	offset := (page - 1) * limit

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		sql = sql + "  ORDER BY pub_date " + sort + " LIMIT ? OFFSET ? "
	} else {
		sql = sql + " ORDER BY pub_date " + sort
	}

	rows, err := ar.db.Query(sql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var articles []models.Article
	var pubDate []byte
	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.ID, &article.Title, &article.Description, &article.Content, &article.Link, &pubDate, &article.Author, &article.SourceID)
		if err != nil {
			return nil, err
		}
		article.PubDate, err = utils.ParseTime(pubDate)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (ar *ArticleRepository) FindOne(id string) (models.Article, error) {
	stmt, err := ar.db.Prepare("SELECT id, title, description, content, link, pub_date, author, source_id FROM articles WHERE id = ?")
	if err != nil {
		return models.Article{}, err
	}
	defer stmt.Close()

	var article models.Article
	var pubDate []byte
	err = stmt.QueryRow(id).Scan(
		&article.ID, &article.Title, &article.Description,
		&article.Content, &article.Link, &pubDate, &article.Author, &article.SourceID)
	if err != nil {
		return models.Article{}, err
	}

	article.PubDate, err = utils.ParseTime(pubDate)
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func (ar *ArticleRepository) Create(article models.Article) (models.Article, error) {
	stmt, err := ar.db.Prepare("INSERT INTO articles (id, title, description, content, link, pub_date, author, source_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return models.Article{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(article.ID, article.Title, article.Description, article.Content, article.Link, &article.PubDate, article.Author, article.SourceID)
	if err != nil {
		return models.Article{}, err
	}

	return models.Article{
		ID:          article.ID,
		Title:       article.Title,
		Description: article.Description,
		Link:        article.Link,
		PubDate:     article.PubDate,
		Author:      article.Author,
		SourceID:    article.SourceID,
	}, nil
}

func (ar *ArticleRepository) Update(id string, article models.Article) (models.Article, error) {
	stmt, err := ar.db.Prepare("UPDATE articles set title = ?, description = ?, content = ?, link = ?, pub_date = ?, author = ?, source_id = ? WHERE id = ?")
	if err != nil {
		return models.Article{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(article.Title, article.Description, article.Content, article.Link, article.PubDate, article.Author, article.SourceID, id)
	if err != nil {
		return models.Article{}, err
	}

	uptadedArticle, err := ar.FindOne(id)
	if err != nil {
		return models.Article{}, err
	}

	return uptadedArticle, nil
}

func (ar *ArticleRepository) Delete(id string) error {
	stmt, err := ar.db.Prepare("DELETE FROM articles WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (ar *ArticleRepository) FindAllBySourceId(sourceID int) ([]models.Article, error) {
	stmt, err := ar.db.Prepare("SELECT * FROM articles WHERE source_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sourceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	var pubDate []byte
	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.ID, &article.Title, &article.Description, &article.Content, &article.Link, &pubDate, &article.Author, &article.SourceID)
		if err != nil {
			return nil, err
		}
		article.PubDate, err = utils.ParseTime(pubDate)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
