package email

import "gopkg.in/gomail.v2"

type Message struct {
	ToEmail          string
	Subject          string
	PlainTextContent string
	HtmlContent      string
}

type EmailSevice struct {
	dialer *gomail.Dialer
	from   string
}

func NewEmailService(smtpHost string, smtpPort int, smtpUser, smtpPassword, fromEmail string) *EmailSevice {
	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)
	return &EmailSevice{
		dialer: dialer,
		from:   fromEmail,
	}
}

func (s *EmailSevice) Send(messageToSend Message) error {
	message := gomail.NewMessage()
	message.SetHeader("From", s.from)
	message.SetHeader("To", messageToSend.ToEmail)
	message.SetHeader("Subject", messageToSend.Subject)
	message.SetBody("text/plain", messageToSend.PlainTextContent)
	message.AddAlternative("text/html", messageToSend.HtmlContent)
	return s.dialer.DialAndSend(message)
}
