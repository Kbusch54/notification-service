package email

import (
	"crypto/tls"
	"fmt"

	"github.com/Kbusch54/notification-service/config"
	"github.com/Kbusch54/notification-service/logg"
	"github.com/Kbusch54/notification-service/notification"
	gomail "gopkg.in/mail.v2"
)

type Service interface {
	SendPriceAlertEmail(notification.NotificationResponse) error
}

type ServiceDefaultImpl struct {
	key      string
	from     string
	smptHost string
	smptPort int
	siteHost string
	log      logg.Logger
}

func NewNotificationService(cfg *config.Config) Service {
	log := logg.NewDefaultLog()
	key := cfg.Services.Brevo.APIKey
	from := cfg.Services.Brevo.Email
	smptHost := cfg.Services.Brevo.Host
	smptPort := cfg.Services.Brevo.Port
	sietHost := cfg.Server.Host
	return &ServiceDefaultImpl{
		log:      log,
		key:      key,
		from:     from,
		smptHost: smptHost,
		smptPort: smptPort,
		siteHost: sietHost,
	}
}

func (s *ServiceDefaultImpl) SendPriceAlertEmail(notification notification.NotificationResponse) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", s.from)
	msg.SetHeader("To", notification.MethodValue.Email)
	msg.SetHeader("Subject", "Price Alert!")
	priceToWatch := notification.TargetPrice
	currentPrice := notification.CurrentPrice
	symbol := notification.Symbol
	msg.SetBody("text/html", fmt.Sprintf("<b>Alert</b> <br><p> This is a test email </p>Price alert for %s <br> Current Price: $%.2f <br> Target Price: $%.2f <br><a href='%s/myinvestments'>View My Investments</a>", symbol, currentPrice, priceToWatch, s.siteHost))
	dialer := gomail.NewDialer(s.smptHost, s.smptPort, s.from, s.key)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := dialer.DialAndSend(msg); err != nil {
		s.log.Errorf("Error sending email: %v", err)
		return err
	}
	s.log.Info("Email sent")
	return nil
}
