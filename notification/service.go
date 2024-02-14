package notification

import (
	"github.com/Kbusch54/notification-service/logg"
	"github.com/Kbusch54/notification-service/persistence/mongodb"
)

type Service interface {
}

type ServiceDefaultImpl struct {
	repo Repository
	log  logg.Logger
}

func NewNotificationService(conn *mongodb.MongoConnection) Service {
	log := logg.NewDefaultLog()
	repos := NewMongoRepository(conn)
	return &ServiceDefaultImpl{
		repo: repos,
		log:  log,
	}
}
