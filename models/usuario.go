package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de usuario de la Base de Datos*/
type Usuario struct {
	// slice de bits
	ID              primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellido        string             `bson:"apellido" json:"apellido,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimieno" json:"fechaNacimieno,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"banner" json:"biografia,omitempty"`
	Ubicacion       string             `bson:"banner" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"banner" json:"sitioWeb,omitempty"`
}
