package bd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/josecardozo13/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* ConsultoRelacion conculta la relacion entre 2 usuarios*/
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("gotwitter")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}
	var resultado models.Relacion
	fmt.Println(resultado)

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err.Error())
		return false, err
	}
	return true, nil
}
