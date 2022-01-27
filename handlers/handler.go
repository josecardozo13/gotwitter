package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/josecardozo13/gotwitter/middlew"
	"github.com/josecardozo13/gotwitter/routers"
	"github.com/rs/cors"
)

/*Manejadores setem mi puerto, el Handler y pongo a escuchar al Servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.BuscoPerfil))).Methods("GET")

	/*PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}*/

	/*handler := cors.AllowAll().Handler(router)*/
	loadServer(router)
	/*log.Fatal(http.ListenAndServe(":"+PORT, handler))*/
}

func loadServer(r *mux.Router) {
	handler := setupCors(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Println(err)
		panic(err)
	}
}

func setupCors(r *mux.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8000",
			"http://localhost:3000",
			"http://localhost:8080",
			"http://localhost:8081",
			"http://sils.com.ar",
			"https://sils.com.ar",
			"http://*.sils.com.ar",
			"https://*.sils.com.ar",
		},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "X-Session-Info"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
	})
	return c.Handler(r)
}
