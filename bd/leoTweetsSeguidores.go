package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/josecardozo13/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* LeoTweersSeguidores lee los tweets de los seguidores relacionado a un usuario*/
func LeoTweersSeguidores(ID string, pagina int) ([]*models.DevuelvoTweetSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("gotwitter")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20
	limit := 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localfield":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}}) // 1 de menor a mayor, -1 de mayor a menor
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": limit})

	cursor, err := col.Aggregate(ctx, condiciones)
	var result []*models.DevuelvoTweetSeguidores

	err = cursor.All(ctx, result)
	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}
	return result, true
}
