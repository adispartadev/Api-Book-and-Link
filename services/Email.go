package services

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(to string, subject string, message string) error {
	body := "From: " + os.Getenv("MAIL_SENDER_NAME") + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", os.Getenv("MAIL_AUTH_EMAIL"), os.Getenv("MAIL_AUTH_PASSWORD"), os.Getenv("MAIL_SMTP_HOST"))
	smtpAddr := fmt.Sprintf("%s:%s", os.Getenv("MAIL_SMTP_HOST"), os.Getenv("MAIL_SMTP_PORT"))

	err := smtp.SendMail(smtpAddr, auth, os.Getenv("MAIL_AUTH_EMAIL"), []string{to}, []byte(body))
	if err != nil {
		return err
	}

	return nil
}
