package notification

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	PRICEALERT = "pricealert"
)

const (
	EMAIL    = "email"
	TELEGRAM = "telegram"
	TWITTER  = "twitter"
)

type Notification struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	UserId       string             `bson:"user_id"`
	NotiType     string             `bson:"noti_type"`
	Symbol       string             `bson:"symbol"`
	PriceToWatch float64            `bson:"price_to_watch"`
	Method       []string           `bson:"method"`
	GreaterThan  bool               `bson:"greater_than"`
	CreatedAt    time.Time          `bson:"created_at"`
	NotifiedAt   time.Time          `bson:"notified_at"`
	Notified     bool               `bson:"notified"`
}

// type MethodValue struct {
// 	Email    string `bson:"email"`
// 	Telegram string `bson:"telegram"`
// 	Twitter  string `bson:"twitter"`
// }

func (s *Notification) GetID() any {
	return s.ID
}

func (s *Notification) toResponse() *NotificationResponse {
	return &NotificationResponse{
		Name:        s.Name,
		Symbol:      s.Symbol,
		Price:       s.PriceToWatch,
		Time:        s.CreatedAt.String(),
		Methods:     s.Method,
		GreaterThan: s.GreaterThan,
		NotifiedAt:  s.NotifiedAt.String(),
		Notified:    s.Notified,
	}
}

// func (s *MethodValue) toResponse() MethodValueResponse {
// 	return MethodValueResponse{
// 		Email:    s.Email,
// 		Telegram: s.Telegram,
// 		Twitter:  s.Twitter,
// 	}
// }

func (s *Notification) NewNotification(name, symbol, userId, notiType string, methods []string, price float64, greaterThan bool) *Notification {
	return &Notification{
		ID:           primitive.NewObjectID(),
		Name:         name,
		UserId:       userId,
		NotiType:     notiType,
		Symbol:       symbol,
		PriceToWatch: price,
		Method:       methods,
		GreaterThan:  greaterThan,
		CreatedAt:    time.Now(),
		NotifiedAt:   time.Time{},
		Notified:     false,
	}
}
