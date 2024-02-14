package notification

import (
	"github.com/Kbusch54/notification-service/logg"
	"github.com/Kbusch54/notification-service/persistence"
	"github.com/Kbusch54/notification-service/persistence/mongodb"
)

const COLLECTION = "notification"

type MongoRepository struct {
	conn    *mongodb.MongoConnection
	absrepo *persistence.AbstractMongoRepository[*Notification]
	log     logg.Logger
}
type Repository interface {
}

func NewMongoRepository(conn *mongodb.MongoConnection) Repository {
	log := logg.NewDefaultLog()
	absrepo := persistence.NewAbstractRepository[*Notification](conn, COLLECTION)
	return &MongoRepository{
		conn:    conn,
		log:     log,
		absrepo: absrepo,
	}
}
