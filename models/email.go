package models

import (
	"github.com/go-mail/mail/v2"
)

const (
	DefaultSender = "support@lenslocked.com"
)

type STMPConfig struct {
	Host string
	Port int 
	Username string 
	Password string
}

func NewEmailService(config STMPConfig) *EmailService {
	es := EmailService{
		dialer: mail.NewDialer(config.Host, config.Port, config.Username, config.Password),
	}
	return &es
}

type EmailService struct {
	// DefaultSender is used as the default sender when one isn't provided for an
	// email. this is also used in functionne where the email is a predetermined,
	// like the forgotten password email
	DefaultSender string

	// un exported fields
	dialer *mail.Dialer
}

