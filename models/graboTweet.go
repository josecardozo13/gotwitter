package models

import (
	"time"
)

/*GraboTweet es la estructura que tendra el tweet*/
type GraboTweet struct {
	UserId  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
