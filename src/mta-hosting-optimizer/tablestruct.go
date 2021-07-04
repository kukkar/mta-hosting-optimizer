package sendinbluetask

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoIpStatusCollection struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	IP       string             `bson:"ip"`
	HostName string             `bson:"hostname"`
	Active   bool               `bson:"active"`
}

type MysqlIpStatusTable struct {
	IP       string `gorm:"ip"`
	HostName string `gorm:"hostname"`
	Active   bool   `gorm:"active"`
}
