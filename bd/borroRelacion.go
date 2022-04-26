package bd

import (
	"context"
	"time"

	"github.com/josecardozo13/gotwitter/models"
)

/*BorroRelacion borra la relacion en la DB*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("gotwitter")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, err
}
