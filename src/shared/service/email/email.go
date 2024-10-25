package email

import "net/smtp"

type EmailService interface {
	Send(recipient, message string) error
}

type emailService struct {
	user     string
	password string
	host     string
}

func NewEmailService(user, password, host string) EmailService {

	return &emailService{
		user:     user,
		password: password,
		host:     host,
	}
}

func (es *emailService) Send(recipient, message string) error {
	auth := smtp.PlainAuth(
		"",
		es.user,
		es.password,
		es.host,
	)

	if err := smtp.SendMail(
		es.host+":587",
		auth,
		es.user,
		[]string{recipient},
		[]byte(message),
	); err != nil {
		return err
	}

	return nil
}
