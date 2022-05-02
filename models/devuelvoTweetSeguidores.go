package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* DevuelvoTweetSeguidores es la estructura con la que devolveremos los twets*/
type DevuelvoTweetSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"userid" json:"userid,omitempty"`
	UsuarioRelacionID string             `bson:"userid" json:"userRelationId,omitempty"`
	Tweet             struct {
		ID      string    `bson:"_id" json:"_id,omitempty"`
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
	}
}
