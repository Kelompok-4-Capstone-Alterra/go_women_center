package helper

import (
	"log"

	"gopkg.in/gomail.v2"
)

type emailSender struct {
	CONFIG_SMTP_HOST     string
	CONFIG_SMTP_PORT     int
	CONFIG_AUTH_EMAIL    string
	CONFIG_AUTH_PASSWORD string
	CONFIG_SENDER_NAME   string
}

func NewEmailSender(SMTP_HOST string, SMTP_PORT int, AUTH_EMAIL string, AUTH_PASSWORD string, SENDER_NAME string) emailSender {
	return emailSender{
		CONFIG_SMTP_HOST:     SMTP_HOST,
		CONFIG_SMTP_PORT:     SMTP_PORT,
		CONFIG_AUTH_EMAIL:    AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD: AUTH_PASSWORD,
		CONFIG_SENDER_NAME:   SENDER_NAME,
	}
}

/*
send an email with html format message body, return nil if success
*/
func (e *emailSender) SendEmail(sendTo, subject, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", e.CONFIG_SENDER_NAME)
	mailer.SetHeader("To", sendTo)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		e.CONFIG_SMTP_HOST,
		e.CONFIG_SMTP_PORT,
		e.CONFIG_AUTH_EMAIL,
		e.CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	log.Println("Mail sent!")
	return nil
}
