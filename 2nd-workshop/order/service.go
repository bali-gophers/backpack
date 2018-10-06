package main

import (
	"crypto/tls"

	gomail "gopkg.in/gomail.v2"
)

type EmailService struct {
	emailAddress string
	dialer       *gomail.Dialer
}

func NewEmailService(user, password string) EmailService {
	dialer := gomail.NewDialer("smtp.gmail.com", 587, user, password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return EmailService{
		emailAddress: user,
		dialer:       dialer,
	}
}

func (service EmailService) Send(email string, subject string, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", service.emailAddress)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	return service.dialer.DialAndSend(msg)
}
