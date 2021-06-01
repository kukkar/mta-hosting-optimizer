package tigerhall

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type MongoTigerCollection struct {
	Id                  bson.ObjectId     `bson:"_id,omitempty"`
	Name                string            `bson:"name"`
	DOB                 time.Time         `bson:"dob"`
	LastSeenAt          time.Time         `bson:"lastSeenAt"`
	LastSeenCoordinates Coordinates       `bson:"lastSeenCoordinates"`
	TigerLastSeenSights []MongoTigerSight `bson:"tigerLastLocations,omitempty"`
}

type MongoTigerSight struct {
	Coordinates Coordinates `bson:"coordinates"`
	TimeStamp   time.Time   `bson:"timeStamp"`
	ImagePath   string      `bson:"image"`
}

type MongoTigerCoordinates struct {
	Lat  float64 `bson:"lat"`
	Long float64 `bson:"long"`
}
