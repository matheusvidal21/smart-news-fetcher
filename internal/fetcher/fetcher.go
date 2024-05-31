package fetcher

import (
	"github.com/google/logger"
	"github.com/matheusvidal21/smart-news-fetcher/internal/articles"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/matheusvidal21/smart-news-fetcher/internal/sources"
	"github.com/mmcdole/gofeed"
	"time"
)

var FeedURL = []string{
	"https://feeds.folha.uol.com.br/ciencia/rss091.xml",
	"https://feeds.folha.uol.com.br/esporte/rss091.xml",
}

type Fetcher struct {
	sourceService  sources.SourceServiceInterface
	articleService articles.ArticleServiceInterface
}

func NewFetcher(sourceService sources.SourceServiceInterface, articleService articles.ArticleServiceInterface) *Fetcher {
	return &Fetcher{sourceService: sourceService, articleService: articleService}
}

func (f *Fetcher) FetchFeeds(url string) {
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(url)
	if err != nil {
		logger.Fatalf("Failed to parse feed: %v", err)
		return
	}

	source := dto.CreateSourceInput{
		Name: feed.Title,
		Url:  url,
	}

	savedSource, err := f.sourceService.Create(source)
	if err != nil {
		logger.Fatalf("Failed to create source: %v", err)
		return
	}
	logger.Info("Source created")

	for _, item := range feed.Items {
		var authorName string
		if len(item.Authors) > 0 {
			authorName = item.Authors[0].Name
		} else {
			authorName = "Unknown"
		}

		article := dto.CreateArticleInput{
			Title:       item.Title,
			Description: item.Description,
			Content:     item.Content,
			Link:        item.Link,
			PubDate:     *item.PublishedParsed,
			Author:      authorName,
			SourceID:    savedSource.ID,
		}
		_, err = f.articleService.Create(article)
		if err != nil {
			logger.Fatalf("Failed to create article: %v", err)
			continue
		}
		logger.Info("Article created: " + item.Title)
	}
}

func (f *Fetcher) StartScheduler(interval time.Duration, url string) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			f.FetchFeeds(url)
		}
	}()
}
