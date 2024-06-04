package service

import (
	"errors"
	"github.com/google/logger"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/matheusvidal21/smart-news-fetcher/internal/email"
	"github.com/matheusvidal21/smart-news-fetcher/internal/interfaces"
	"github.com/matheusvidal21/smart-news-fetcher/internal/models"
	"time"
)

type SourceService struct {
	sourceRepository interfaces.SourceRepositoryInterface
	userService      interfaces.UserServiceInterface
	emailService     interfaces.EmailService
	fetcher          interfaces.FetcherInterface
	quitChannels     map[int]chan struct{}
}

func NewSourceService(sourceRepository interfaces.SourceRepositoryInterface, userService interfaces.UserServiceInterface, emailService interfaces.EmailService, fetcher interfaces.FetcherInterface) *SourceService {
	return &SourceService{
		sourceRepository: sourceRepository,
		userService:      userService,
		emailService:     emailService,
		fetcher:          fetcher,
		quitChannels:     make(map[int]chan struct{}),
	}
}

func (sr *SourceService) FindAll(page, limit int, sort string) ([]models.Source, error) {
	sources, err := sr.sourceRepository.FindAll(page, limit, sort)
	if err != nil {
		return []models.Source{}, errors.New("Failed to find sources: " + err.Error())
	}
	return sources, nil
}

func (sr *SourceService) FindOne(id int) (dto.FindOneSourceOutput, error) {
	source, err := sr.sourceRepository.FindOne(id)

	if err != nil {
		return dto.FindOneSourceOutput{}, errors.New("Failed to find source: " + err.Error())
	}
	return dto.FindOneSourceOutput{
		ID:             source.ID,
		Name:           source.Name,
		Url:            source.Url,
		UpdateInterval: source.UpdateInterval,
		UserID:         source.UserID,
		SavedAt:        source.SavedAt,
	}, nil
}

func (sr *SourceService) Create(sourceDto dto.CreateSourceInput) (dto.CreateSourceOutput, error) {
	user, err := sr.userService.FindById(sourceDto.UserID)
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

	source := models.Source{
		Name:           sourceDto.Name,
		Url:            sourceDto.Url,
		UserID:         sourceDto.UserID,
		UpdateInterval: sourceDto.UpdateInterval,
		SavedAt:        time.Now(),
	}

	sourceSaved, err := sr.sourceRepository.Create(source)
	if err != nil {
		return dto.CreateSourceOutput{}, errors.New("Failed to save source: " + err.Error())
	}

	message := email.Message{
		ToEmail:          user.Email,
		Subject:          "Smart News Fetcher - New Source",
		PlainTextContent: "New source has been added: " + sourceDto.Url,
		HtmlContent:      "<p>New source has been added: <a>" + sourceDto.Url + "</a></p>",
	}

	err = sr.emailService.Send(message)
	if err != nil {
		return dto.CreateSourceOutput{}, errors.New("Failed to send email: " + err.Error())
	}

	sr.fetcher.StoreFeed(sourceSaved.ID, feed)

	return dto.CreateSourceOutput{
		ID:             sourceSaved.ID,
		Name:           sourceSaved.Name,
		Url:            sourceSaved.Url,
		UpdateInterval: sourceSaved.UpdateInterval,
		UserID:         sourceSaved.UserID,
		SavedAt:        sourceSaved.SavedAt,
	}, nil
}

func (sr *SourceService) Update(id int, sourceDto dto.UpdateSourceInput) (dto.UpdateSourceOutput, error) {
	currentSource, err := sr.sourceRepository.FindOne(id)
	if err != nil {
		return dto.UpdateSourceOutput{}, errors.New("Failed to find source: " + err.Error())
	}

	if !(currentSource.Url == sourceDto.Url) {
		feed, err := sr.fetcher.ParseFeed(sourceDto.Url)
		if err != nil {
			return dto.UpdateSourceOutput{}, errors.New("Failed to parse source feed: " + err.Error())
		}
		sr.fetcher.StoreFeed(id, feed)
	}

	source := models.Source{
		Name:           sourceDto.Name,
		UpdateInterval: sourceDto.UpdateInterval,
		Url:            sourceDto.Url,
	}

	sourceUpdated, err := sr.sourceRepository.Update(id, source)
	if err != nil {
		return dto.UpdateSourceOutput{}, errors.New("Failed to update source: " + err.Error())
	}
	return dto.UpdateSourceOutput{
		ID:             sourceUpdated.ID,
		Name:           sourceUpdated.Name,
		Url:            sourceUpdated.Url,
		UserID:         sourceUpdated.UserID,
		UpdateInterval: sourceUpdated.UpdateInterval,
		SavedAt:        sourceUpdated.SavedAt,
	}, err
}

func (sr *SourceService) Delete(id int) error {
	err := sr.sourceRepository.Delete(id)
	if err != nil {
		return errors.New("Failed to delete source: " + err.Error())
	}
	return nil
}

func (sr *SourceService) LoadFeed(id int) error {
	source, err := sr.sourceRepository.FindOne(id)
	if err != nil {
		return errors.New("Source not found: " + err.Error())
	}

	feed, err := sr.fetcher.LoadFeed(source.ID)
	if err != nil {
		feed, err = sr.fetcher.ParseFeed(source.Url)
		if err != nil {
			return errors.New("Failed to load feed: " + err.Error())
		}

		sr.fetcher.StoreFeed(source.ID, feed)
	}

	sr.fetcher.StoreFeed(source.ID, feed)
	sr.fetcher.StartScheduler(source, feed)
	return nil
}

func (sr *SourceService) FindByUserId(userId int) ([]models.Source, error) {
	sources, err := sr.sourceRepository.FindByUserId(userId)
	if err != nil {
		return []models.Source{}, errors.New("Failed to find sources: " + err.Error())
	}
	return sources, nil
}

func (sr *SourceService) SubscribeToNewsletter(id int) error {
	source, err := sr.sourceRepository.FindOne(id)
	if err != nil {
		return errors.New("Source not found: " + err.Error())
	}

	if source.Subscriber {
		if _, exists := sr.quitChannels[id]; !exists {
			sr.StartSubscription(source)
		}
		return errors.New("Source already subscribed")
	}

	source.Subscriber = true

	_, err = sr.sourceRepository.Update(id, source)
	if err != nil {
		return errors.New("Failed to update source: " + err.Error())
	}

	sr.StartSubscription(source)
	return nil
}

func (sr *SourceService) UnsubscribeFromNewsletter(id int) error {
	source, err := sr.sourceRepository.FindOne(id)
	if err != nil {
		return errors.New("Source not found: " + err.Error())
	}

	if source.Subscriber == false {
		return errors.New("Source already unsubscribed")
	}

	source.Subscriber = false
	_, err = sr.sourceRepository.Update(id, source)
	if err != nil {
		return errors.New("Failed to update source: " + err.Error())
	}

	if quit, exists := sr.quitChannels[id]; exists {
		sr.quitChannels[id] <- struct{}{}
		close(quit)
		delete(sr.quitChannels, id)
	}
	return nil
}

func (sr *SourceService) StartSubscription(source models.Source) {
	ticker := time.NewTicker(24 * time.Hour)
	sr.quitChannels[source.ID] = make(chan struct{})

	go func() {
		user, _ := sr.userService.FindById(source.UserID)
		for {
			select {
			case <-ticker.C:
				feed, err := sr.fetcher.ParseFeed(source.Url)
				if err != nil {
					logger.Errorf("Failed to parse source feed: %v", err)
					continue
				}
				sr.fetcher.StoreFeed(source.ID, feed)

				textContent := "Current articles on " + source.Name + ":\n"
				htmlContent := "<p>Current articles on " + source.Name + ":</p>"

				for _, item := range feed.Items {
					textContent += "- " + item.Title + "\n"
					textContent += item.Description + "\n"
					textContent += item.Link + "\n\n"
					textContent += "------------------------------------------------------\n\n"

					htmlContent += "<p><b>" + item.Title + "</b><br>"
					htmlContent += item.Description + "<br>"
					htmlContent += "<hr>"
				}

				message := email.Message{
					ToEmail:          user.Email,
					Subject:          "Newsletter: " + source.Name + " - " + "Day " + time.Now().Format("2006-01-02") + "!",
					PlainTextContent: textContent,
					HtmlContent:      htmlContent,
				}

				err = sr.emailService.Send(message)
				if err != nil {
					logger.Errorf("Failed to send email: %v", err)
					continue
				}

			case <-sr.quitChannels[source.ID]:
				ticker.Stop()
				return
			}
		}
	}()
}

func (sr *SourceService) InitializeSubscription() {
	sources, err := sr.sourceRepository.FindAllActive()
	if err != nil {
		logger.Errorf("Failed to initialize subscriptions: %v", err)
		return
	}
	for _, source := range sources {
		sr.StartSubscription(source)
	}
}
