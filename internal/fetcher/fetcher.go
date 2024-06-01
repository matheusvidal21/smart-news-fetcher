package fetcher

import (
	"github.com/google/logger"
	"github.com/matheusvidal21/smart-news-fetcher/internal/articles"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/mmcdole/gofeed"
	"sync"
	"time"
)

type FetcherInterface interface {
	GetFeedChannel(url string) chan *gofeed.Feed
	ParseFeed(url string) (*gofeed.Feed, error)
	FetchFeeds(id int, feed *gofeed.Feed)
	StartScheduler(interval time.Duration, feed *gofeed.Feed, id int)
}

type Fetcher struct {
	articleService articles.ArticleServiceInterface
	feedChannels   map[string]chan *gofeed.Feed
	mu             sync.Mutex
}

func NewFetcher(articleService articles.ArticleServiceInterface) *Fetcher {
	return &Fetcher{
		articleService: articleService,
		feedChannels:   make(map[string]chan *gofeed.Feed),
	}
}

func (f *Fetcher) GetFeedChannel(url string) chan *gofeed.Feed {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ch, ok := f.feedChannels[url]; ok {
		return ch
	}
	ch := make(chan *gofeed.Feed, 1)
	f.feedChannels[url] = ch
	return ch
}

func (f *Fetcher) ParseFeed(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(url)
	if err != nil {
		logger.Errorf("Failed to parse feed: %v", err)
		return nil, err
	}
	logger.Info("Feed parsed: " + url)
	return feed, nil
}

func (f *Fetcher) FetchFeeds(id int, feed *gofeed.Feed) {
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
			SourceID:    id,
		}
		_, err := f.articleService.Create(article)
		if err != nil {
			logger.Errorf("Failed to create article: %v", err)
			continue
		}
		logger.Info("Article created: " + item.Title)
	}
}

func (f *Fetcher) StartScheduler(interval time.Duration, feed *gofeed.Feed, id int) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			f.FetchFeeds(id, feed)
		}
	}()
}
