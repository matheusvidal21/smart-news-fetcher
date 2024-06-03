package fetcher

import (
	"errors"
	"github.com/google/logger"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/matheusvidal21/smart-news-fetcher/internal/interfaces"
	"github.com/matheusvidal21/smart-news-fetcher/internal/models"
	"github.com/mmcdole/gofeed"
	"github.com/patrickmn/go-cache"
	"strconv"
	"time"
)

type Fetcher struct {
	articleService interfaces.ArticleServiceInterface
	cache          *cache.Cache
}

func NewFetcher(articleService interfaces.ArticleServiceInterface) *Fetcher {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &Fetcher{
		articleService: articleService,
		cache:          c,
	}
}

func (f *Fetcher) LoadFeed(id int) (*gofeed.Feed, error) {
	feedKey := "feed_" + strconv.Itoa(id)
	if cacheFeed, found := f.cache.Get(feedKey); found {
		logger.Info("Feed found in cache: " + feedKey)
		return cacheFeed.(*gofeed.Feed), nil
	}
	logger.Info("Feed not found in cache: " + feedKey)
	return nil, errors.New("Feed not found in cache: " + feedKey)
}

func (f *Fetcher) StoreFeed(id int, feed *gofeed.Feed) {
	feedKey := "feed_" + strconv.Itoa(id)
	f.cache.Set(feedKey, feed, cache.DefaultExpiration)
	logger.Info("Feed stored in cache: " + feedKey)
}

func (f *Fetcher) ParseFeed(url string) (*gofeed.Feed, error) {
	if cachedFeed, found := f.cache.Get(url); found {
		logger.Info("Feed loaded from cache: " + url)
		return cachedFeed.(*gofeed.Feed), nil
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		logger.Errorf("Failed to parse feed: %v", err)
		return nil, err
	}

	f.cache.Set(url, feed, cache.DefaultExpiration)
	logger.Info("Feed parsed and cached: " + url)
	return feed, nil
}

func (f *Fetcher) FetchFeeds(id, limit int, feed *gofeed.Feed) {
	count := 0
	for _, item := range feed.Items {
		if limit > 0 && count >= limit {
			break
		}
		count++

		var authorName string
		if len(item.Authors) > 0 {
			authorName = item.Authors[0].Name
		} else {
			authorName = "Unknown"
		}

		guid, _ := strconv.Atoi(item.GUID)
		article := dto.CreateArticleInput{
			ID:          guid,
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

func (f *Fetcher) StartScheduler(source models.Source, feed *gofeed.Feed) {
	interval := time.Duration(source.UpdateInterval) * time.Minute
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			f.FetchFeeds(source.ID, -1, feed)
		}
	}()
}
