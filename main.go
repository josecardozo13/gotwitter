package main

import (
	"log"

	"github.com/josecardozo13/gotwitter/bd"
	"github.com/josecardozo13/gotwitter/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("sin conexion a la BD")
	}
	handlers.Manejadores()
}
