package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoCN       = ConectarDB()                                                                                         // Variable a exportar a toda la aplicacion
	clientOptions = options.Client().ApplyURI("mongodb+srv://dbtwitter123:dbtwitter123@cluster0.3flyr.mongodb.net/test") // Variable interna
)

// ConectarDB es la funcion que mepermite conectarm a la BD
func ConectarDB() *mongo.Client {
	// el context es como si fuera un espacio en memoria para compartir cosas, tb puedo crear variables mientras dure esa operacion rutina
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil) // es diferente al anterior
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la BD")
	return client
}

// CheckConnection es el ping a la BD
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
