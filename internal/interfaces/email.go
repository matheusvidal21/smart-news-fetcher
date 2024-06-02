package interfaces

import "github.com/matheusvidal21/smart-news-fetcher/internal/email"

type EmailService interface {
	Send(messageToSend email.Message) error
}
