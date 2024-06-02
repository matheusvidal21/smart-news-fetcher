package sources

import (
	"errors"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/matheusvidal21/smart-news-fetcher/internal/fetcher"
	"github.com/matheusvidal21/smart-news-fetcher/internal/user"
	"time"
)

type SourceServiceInterface interface {
	FindAll(page, limit int, sort string) ([]Source, error)
	FindOne(id int) (dto.FindOneSourceOutput, error)
	Create(sourceDto dto.CreateSourceInput) (dto.CreateSourceOutput, error)
	Update(id int, sourceDto dto.UpdateSourceInput) (dto.UpdateSourceOutput, error)
	Delete(id int) error
	LoadFeed(id int, duration time.Duration) error
	FindByUserId(userId int) ([]Source, error)
}

type SourceService struct {
	sourceRepository SourceRepositoryInterface
	userService      user.UserServiceInterface
	fetcher          fetcher.FetcherInterface
}

func NewSourceService(sourceRepository SourceRepositoryInterface, userService user.UserServiceInterface, fetcher fetcher.FetcherInterface) *SourceService {
	return &SourceService{
		sourceRepository: sourceRepository,
		userService:      userService,
		fetcher:          fetcher,
	}
}

func (sr *SourceService) FindAll(page, limit int, sort string) ([]Source, error) {
	sources, err := sr.sourceRepository.FindAll(page, limit, sort)
	if err != nil {
		return []Source{}, errors.New("Failed to find sources: " + err.Error())
	}
	return sources, nil
}

func (sr *SourceService) FindOne(id int) (dto.FindOneSourceOutput, error) {
	source, err := sr.sourceRepository.FindOne(id)

	if err != nil {
		return dto.FindOneSourceOutput{}, errors.New("Failed to find source: " + err.Error())
	}
	return dto.FindOneSourceOutput{
		ID:      source.ID,
		Name:    source.Name,
		Url:     source.Url,
		UserID:  source.UserID,
		SavedAt: source.SavedAt,
	}, nil
}

func (sr *SourceService) Create(sourceDto dto.CreateSourceInput) (dto.CreateSourceOutput, error) {
	_, err := sr.userService.FindById(sourceDto.UserID)
	if err != nil {
		return dto.CreateSourceOutput{}, errors.New("User not found: " + err.Error())
	}

	sources, _ := sr.sourceRepository.FindByUserId(sourceDto.UserID)
	for _, source := range sources {
		if source.Url == sourceDto.Url {
			return dto.CreateSourceOutput{}, errors.New("Source already exists: " + sourceDto.Url)
		}
	}

	feed, err := sr.fetcher.ParseFeed(sourceDto.Url)
	if err != nil {
		return dto.CreateSourceOutput{}, errors.New("Failed to parse source feed: " + err.Error())
	}

	feedCh := sr.fetcher.GetFeedChannel(sourceDto.Url)
	feedCh <- feed

	source := Source{
		Name:    sourceDto.Name,
		Url:     sourceDto.Url,
		UserID:  sourceDto.UserID,
		SavedAt: time.Now(),
	}

	sourceSaved, err := sr.sourceRepository.Create(source)
	if err != nil {
		return dto.CreateSourceOutput{}, errors.New("Failed to save source: " + err.Error())
	}

	return dto.CreateSourceOutput{
		ID:      sourceSaved.ID,
		Name:    sourceSaved.Name,
		Url:     sourceSaved.Url,
		UserID:  sourceSaved.UserID,
		SavedAt: sourceSaved.SavedAt,
	}, nil
}

func (sr *SourceService) Update(id int, sourceDto dto.UpdateSourceInput) (dto.UpdateSourceOutput, error) {
	feed, err := sr.fetcher.ParseFeed(sourceDto.Url)
	if err != nil {
		return dto.UpdateSourceOutput{}, errors.New("Failed to parse source feed: " + err.Error())
	}

	feedCh := sr.fetcher.GetFeedChannel(sourceDto.Url)
	feedCh <- feed

	source := Source{
		Name: sourceDto.Name,
		Url:  sourceDto.Url,
	}
	sourceUpdated, err := sr.sourceRepository.Update(id, source)
	if err != nil {
		return dto.UpdateSourceOutput{}, errors.New("Failed to update source: " + err.Error())
	}
	return dto.UpdateSourceOutput{
		ID:      sourceUpdated.ID,
		Name:    sourceUpdated.Name,
		Url:     sourceUpdated.Url,
		SavedAt: sourceUpdated.SavedAt,
	}, err
}

func (sr *SourceService) Delete(id int) error {
	err := sr.sourceRepository.Delete(id)
	if err != nil {
		return errors.New("Failed to delete source: " + err.Error())
	}
	return nil
}

func (sr *SourceService) LoadFeed(id int, duration time.Duration) error {
	source, err := sr.sourceRepository.FindOne(id)
	if err != nil {
		return errors.New("Failed to find source: " + err.Error())
	}

	feedCh := sr.fetcher.GetFeedChannel(source.Url)
	feed := <-feedCh
	sr.fetcher.StartScheduler(duration, feed, id)
	return nil
}

func (sr *SourceService) FindByUserId(userId int) ([]Source, error) {
	sources, err := sr.sourceRepository.FindByUserId(userId)
	if err != nil {
		return []Source{}, errors.New("Failed to find sources: " + err.Error())
	}
	return sources, nil
}
