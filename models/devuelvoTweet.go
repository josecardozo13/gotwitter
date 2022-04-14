package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* DevuelvoTweet es la estructura con la que devolveremos los twets*/
type DevuelvoTweet struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId  string             `bson:"userid" json:"userid,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
