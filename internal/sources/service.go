package sources

import (
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/matheusvidal21/smart-news-fetcher/internal/fetcher"
	"time"
)

type SourceServiceInterface interface {
	FindAll(page, limit int, sort string) ([]Source, error)
	FindOne(id int) (dto.FindOneSourceOutput, error)
	Create(sourceDto dto.CreateSourceInput) (dto.CreateSourceOutput, error)
	Update(id int, sourceDto dto.UpdateSourceInput) (dto.UpdateSourceOutput, error)
	Delete(id int) error
	LoadFeed(id int) error
}

type SourceService struct {
	sourceRepository SourceRepositoryInterface
}

func NewSourceService(sourceRepository SourceRepositoryInterface) *SourceService {
	return &SourceService{sourceRepository: sourceRepository}
}

func (sr *SourceService) FindAll(page, limit int, sort string) ([]Source, error) {
	return sr.sourceRepository.FindAll(page, limit, sort)
}

func (sr *SourceService) FindOne(id int) (dto.FindOneSourceOutput, error) {
	source, err := sr.sourceRepository.FindOne(id)

	if err != nil {
		return dto.FindOneSourceOutput{}, err
	}
	return dto.FindOneSourceOutput{
		ID:      source.ID,
		Name:    source.Name,
		Url:     source.Url,
		SavedAt: source.SavedAt,
	}, nil
}

func (sr *SourceService) Create(sourceDto dto.CreateSourceInput) (dto.CreateSourceOutput, error) {
	source := Source{
		Name:    sourceDto.Name,
		Url:     sourceDto.Url,
		SavedAt: time.Now(),
	}
	sourceSaved, err := sr.sourceRepository.Create(source)
	if err != nil {
		return dto.CreateSourceOutput{}, err
	}

	return dto.CreateSourceOutput{
		ID:      sourceSaved.ID,
		Name:    sourceSaved.Name,
		Url:     sourceSaved.Url,
		SavedAt: sourceSaved.SavedAt,
	}, nil
}

func (sr *SourceService) Update(id int, sourceDto dto.UpdateSourceInput) (dto.UpdateSourceOutput, error) {
	source := Source{
		Name: sourceDto.Name,
		Url:  sourceDto.Url,
	}
	sourceUpdated, err := sr.sourceRepository.Update(id, source)
	if err != nil {
		return dto.UpdateSourceOutput{}, err
	}
	return dto.UpdateSourceOutput{
		ID:      sourceUpdated.ID,
		Name:    sourceUpdated.Name,
		Url:     sourceUpdated.Url,
		SavedAt: sourceUpdated.SavedAt,
	}, err
}

func (sr *SourceService) Delete(id int) error {
	return sr.sourceRepository.Delete(id)
}

func (sr *SourceService) LoadFeed(id int, fetcher fetcher.Fetcher) error {
	source, err := sr.FindOne(id)
	if err != nil {
		return err
	}

	fetcher.StartScheduler(time.Minute*10, source.Url)
	return nil
}
