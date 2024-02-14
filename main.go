package main

import (
	"crypto/tls"

	"github.com/Kbusch54/notification-service/config"
	"github.com/Kbusch54/notification-service/logg"
	gomail "gopkg.in/mail.v2"
)

func main() {
	cfg := config.Load("./config/env")

	logger := logg.NewDefaultLog()

	key := cfg.Services.Brevo.APIKey
	from := cfg.Services.Brevo.Email
	smptHost := cfg.Services.Brevo.Host
	smptPort := cfg.Services.Brevo.Port
	to := "kevinbusch54@gmail.com"

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Alert!")
	msg.SetBody("text/html", "<b>Alert</b> \n <p> This is a test email </p>\nPrice alert for Apple Inc. (AAPL) \n Current Price: $100.00 \n Target Price: $110.00 \n")
	dialer := gomail.NewDialer(smptHost, smptPort, from, key)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := dialer.DialAndSend(msg); err != nil {
		logger.Errorf("Error sending email: %v", err)
	}
	logger.Info("Email sent")

}
