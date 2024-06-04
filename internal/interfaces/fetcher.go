package interfaces

import (
	"github.com/matheusvidal21/smart-news-fetcher/internal/models"
	"github.com/mmcdole/gofeed"
)

type FetcherInterface interface {
	LoadFeed(id int) (*gofeed.Feed, error)
	StoreFeed(id int, feed *gofeed.Feed)
	ParseFeed(url string) (*gofeed.Feed, error)
	FetchFeeds(id int, feed *gofeed.Feed)
	StartScheduler(source models.Source, feed *gofeed.Feed)
}
