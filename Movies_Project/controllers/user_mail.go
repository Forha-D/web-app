package controllers

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
	"net/http"
)

const (
	SMTPHost     = "smtp@gmail.com"
	SMTPPort     = 587
	SMTPUsername = "example@gmail.com"
	SMTPPassword = "Password"
	SenderEmail  = "sender@example.com"
)

func sendEmail(toEmail, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", SenderEmail)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(SMTPHost, SMTPPort, SMTPUsername, SMTPPassword)

	// Uncomment the following line if you need to use SSL/TLS
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func SendEmailHandler(c echo.Context) error {
	toEmail := c.FormValue("to_email")
	subject := c.FormValue("subject")
	body := c.FormValue("body")

	if err := sendEmail(toEmail, subject, body); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to send the email")
	}

	return c.String(http.StatusOK, "Email sent successfully!")
}
